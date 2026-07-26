import NewRoomCommand from './command/new-room'
import JoinRoomCommand from './command/join-room'
import { plainToInstance } from 'class-transformer'
import { validateOrReject } from 'class-validator'
import { State } from './state/state'
import ChooseTeamCommand from './command/choose-team'
import { TeamId } from './team-id'
import StartGameCommand from './command/start-game'
import PlayCardMove from './command/move/play-card'
import type { Card, Suit } from './card'
import PlayTurnCommand from './command/play-turn'
import AcceptTrumpMove from './command/move/accept-trump'
import SelectTrumpMove from './command/move/select-trump'
import type { SessionManager } from './session-manager'

enum CloseEventCodes {
	InvalidSession = 4001,
}

export class GameClient {
	private ws: WebSocket | null = null
	private sessionManager: SessionManager
	private wsUrl: string
	private listeners: ((state: State | null) => void)[] = []

	constructor(wsUrl: string, sessionManager: SessionManager) {
		this.wsUrl = wsUrl
		this.sessionManager = sessionManager
	}

	public async restoreConnection(): Promise<void> {
		if (this.ws && this.ws.readyState !== WebSocket.CLOSED) {
			return
		}

		const session = await this.sessionManager.getSession()
		if (!session) {
			this.fireStateUpdate(null)
			return
		}
		await this.connect(session.userId)
	}

	public async connect(userId: string): Promise<void> {
		let session = await this.sessionManager.getSession()
		if (session && session.userId !== userId) {
			await this.sessionManager.clearSession()
			session = null
		}

		const url =
			`${this.wsUrl}?userId=${encodeURIComponent(userId)}` +
			(session ? `&sessionId=${encodeURIComponent(session.id)}` : '')

		return new Promise((resolve, reject) => {
			this.ws = new WebSocket(url)

			this.ws.onopen = () => {
				console.log('WebSocket connection established')
				resolve()
			}

			this.ws.onerror = (err) => {
				console.error('WebSocket error:', err)
				reject(err)
			}

			this.ws.onclose = (event) => {
				console.log('WebSocket closed: ', event.code, event.reason)
				this.ws = null

				if (event.code === CloseEventCodes.InvalidSession) {
					console.warn('Session invalid, clearing session data')
					void this.sessionManager.clearSession()
					this.fireStateUpdate(null)
					return
				}

				console.log('Connection lost')
				setTimeout(() => {
					void this.restoreConnection()
				}, 2000)
			}

			this.ws.onmessage = (event) => {
				if (!event.data) {
					console.warn('Received empty message from WebSocket')
					return
				}
				const content = JSON.parse(event.data)
				if (content.sessionId) {
					void this.sessionManager.saveSession(
						content.sessionId,
						userId
					)
					return
				}
				const message = plainToInstance(State, content as State)
				validateOrReject(message)
				this.fireStateUpdate(message)
			}
		})
	}

	public disconnect() {
		void this.sessionManager.clearSession()
		if (this.ws) {
			this.ws.close()
		}
		this.fireStateUpdate(null)
		this.ws = null
	}

	public joinRoom(roomId: string): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const command = new JoinRoomCommand(roomId)
		this.ws.send(JSON.stringify(command))
	}

	public chooseTeam(teamId: TeamId): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		if (teamId === TeamId.NoTeam) {
			throw new Error('Cannot choose NoTeam.')
		}

		const command = new ChooseTeamCommand(teamId)
		this.ws.send(JSON.stringify(command))
	}

	public createRoom(): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const command = new NewRoomCommand()
		this.ws.send(JSON.stringify(command))
	}

	public startGame(): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const command = new StartGameCommand()
		this.ws.send(JSON.stringify(command))
	}

	public playCard(card: Card, skipDeclarations: boolean = false): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const move = new PlayCardMove(card, skipDeclarations)
		const command = new PlayTurnCommand(move)
		this.ws.send(JSON.stringify(command))
	}

	public acceptTrump(accepted: boolean): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const move = new AcceptTrumpMove(accepted)
		const command = new PlayTurnCommand(move)
		this.ws.send(JSON.stringify(command))
	}

	public selectTrump(suit: Suit | null): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.')
		}

		const move = new SelectTrumpMove(suit)
		const command = new PlayTurnCommand(move)
		this.ws.send(JSON.stringify(command))
	}

	public addListener(listener: (state: State | null) => void): void {
		this.listeners.push(listener)
	}

	public removeListener(listener: (state: State | null) => void): void {
		this.listeners = this.listeners.filter((l) => l !== listener)
	}

	private fireStateUpdate(state: State | null): void {
		this.listeners.forEach((listener) => listener(state))
	}
}

export default GameClient
