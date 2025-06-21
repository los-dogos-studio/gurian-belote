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

export class GameClient {
	private ws: WebSocket | null = null;
	private wsUrl: string;
	private listeners: ((state: State) => void)[] = [];

	constructor(wsUrl: string) {
		this.wsUrl = wsUrl;
	}

	public connect(userId: string): Promise<void> {
		return new Promise((resolve, reject) => {
			this.ws = new WebSocket(this.wsUrl + `?userId=${encodeURIComponent(userId)}`);

			this.ws.onopen = () => {
				console.log('WebSocket connection established');
				resolve();
			};

			this.ws.onerror = (err) => {
				console.error('WebSocket error:', err);
				reject(err);
			};

			this.ws.onclose = () => {
				console.log('WebSocket closed');
				this.ws = null;
			};

			this.ws.onmessage = (event) => {
				if (!event.data) {
					console.warn('Received empty message from WebSocket');
					return;
				}

				const message = plainToInstance(State, JSON.parse(event.data) as State);
				validateOrReject(message);
				this.listeners.forEach(listener => listener(message));
			};
		});
	}

	public disconnect() {
		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}
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


	public addListener(listener: (state: State) => void): void {
		this.listeners.push(listener);
	}
}

export default GameClient;
