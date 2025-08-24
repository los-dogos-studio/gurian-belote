import NewRoomCommand from './command/new-room';
import JoinRoomCommand from './command/join-room';
import { plainToInstance } from 'class-transformer';
import { validateOrReject } from 'class-validator';
import { State } from './state/state';
import ChooseTeamCommand from './command/choose-team';
import { TeamId } from './team-id';
import StartGameCommand from './command/start-game';
import PlayCardMove from './command/move/play-card';
import type { Card, Suit } from './card';
import PlayTurnCommand from './command/play-turn';
import AcceptTrumpMove from './command/move/accept-trump';
import SelectTrumpMove from './command/move/select-trump';
import { SessionManager } from './session-manager';

export class GameClient {
	private ws: WebSocket | null = null;
	private wsUrl: string;
	private authUrl: string;
	private listeners: ((state: State | null) => void)[] = [];
	private sessionManager: SessionManager;

	constructor(wsUrl: string, authUrl: string, sessionManager: SessionManager) {
		this.wsUrl = wsUrl;
		this.authUrl = authUrl;
		this.sessionManager = sessionManager;
	}

	public connect(userName: string, callbacks?: { onSuccess?: () => void; onError?: (error: Error) => void }): void {
		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}

		this.authenticate(userName).then(() => {
			const url = `${this.wsUrl}?userName=${encodeURIComponent(userName)}`;
			this.createWebSocketConnection(url, {
				onRoomJoin: (roomId) => this.onSuccessfulRoomJoin(roomId, userName),
				onSuccess: callbacks?.onSuccess,
				onError: callbacks?.onError
			});
		}).catch(callbacks?.onError);
	}

	public reconnect(callbacks?: { onSuccess?: () => void; onError?: (error: Error) => void }): void {
		const stored = this.getStoredSession();
		if (!stored?.roomId) {
			callbacks?.onError?.(new Error('No stored session to reconnect with'));
			return;
		}

		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}

		const url = `${this.wsUrl}?roomId=${stored.roomId}`;
		this.createWebSocketConnection(url, {
			onRoomJoin: (roomId) => this.onSuccessfulRoomJoin(roomId, stored.userName),
			onSuccess: callbacks?.onSuccess,
			onError: callbacks?.onError
		});
	}

	private createWebSocketConnection(url: string, handlers: {
		onRoomJoin: (roomId: string) => void;
		onSuccess?: () => void;
		onError?: (error: Error) => void;
	}): void {
		const ws = new WebSocket(url);

		ws.onopen = () => {
			this.ws = ws;
			handlers.onSuccess?.();
		};

		ws.onerror = (err) => {
			console.error('WebSocket error:', err);
			handlers.onError?.(new Error('WebSocket connection failed'));
		};

		ws.onclose = () => {
			this.ws = null;
		};

		ws.onmessage = async (event) => {
			if (!event.data) {
				return;
			}

			const data = JSON.parse(event.data);

			if (data.error) {
				handlers.onError?.(new Error(data.error));
				return;
			}

			if (!data.gameState) {
				return;
			}

			const message = plainToInstance(State, data as State);
			try {
				await validateOrReject(message);
			} catch(error: any) {
				handlers.onError?.(new Error("Invalid game state received"));
			}

			if (message.gameState?.roomId) { // TODO shouldn't do on every message
				handlers.onRoomJoin(message.gameState.roomId);
			}

			this.listeners.forEach(listener => listener(message));
		};
	}


	private async authenticate(userName: string): Promise<void> {
		const response = await fetch(`${this.authUrl}?userName=${encodeURIComponent(userName)}`, {
			method: 'POST',
			credentials: 'include'
		});

		if (!response.ok) {
			throw new Error('Authentication failed');
		}
	}

	private onSuccessfulRoomJoin(roomId: string, userName: string) {
		this.sessionManager.saveSession({
			roomId,
			userName,
			autoReconnect: true
		});
	}

	public disconnect() {
		if (this.ws) {
			this.ws.close();
		}
		this.ws = null;
		this.listeners = [];
	}

	public leaveRoom() {
		this.sessionManager.setAutoReconnect(false);
		this.listeners.forEach(listener => listener(null));
		if (this.ws) {
			this.ws.close();
		}
		this.ws = null;
	}

	public getStoredSession() {
		return this.sessionManager.getSession();
	}


	public joinRoom(roomId: string): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const command = new JoinRoomCommand(roomId);
		this.ws.send(JSON.stringify(command));
	}

	public chooseTeam(teamId: TeamId): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		if (teamId === TeamId.NoTeam) {
			throw new Error('Cannot choose NoTeam.');
		}

		const command = new ChooseTeamCommand(teamId);
		this.ws.send(JSON.stringify(command));
	}

	public createRoom(): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const command = new NewRoomCommand();
		this.ws.send(JSON.stringify(command));
	}

	public startGame(): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const command = new StartGameCommand();
		this.ws.send(JSON.stringify(command));
	}

	public playCard(card: Card): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const move = new PlayCardMove(card);
		const command = new PlayTurnCommand(move);
		this.ws.send(JSON.stringify(command));
	}

	public acceptTrump(accepted: boolean): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const move = new AcceptTrumpMove(accepted);
		const command = new PlayTurnCommand(move);
		this.ws.send(JSON.stringify(command));
	}

	public selectTrump(suit: Suit | null): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const move = new SelectTrumpMove(suit);
		const command = new PlayTurnCommand(move);
		this.ws.send(JSON.stringify(command));
	}

	public addListener(listener: (state: State | null) => void): void {
		this.listeners.push(listener);
	}
}

export default GameClient;
