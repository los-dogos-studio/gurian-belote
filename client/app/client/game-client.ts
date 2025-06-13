import NewRoomCommand from './command/new-room';
import JoinRoomCommand from './command/join-room';

export class GameClient {
	private ws: WebSocket | null = null;
	private wsUrl: string;

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
				const message = JSON.parse(event.data);
				console.log('Received message:', message);
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

	public createRoom(): void {
		if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
			throw new Error('WebSocket is not connected.');
		}

		const command = new NewRoomCommand();
		this.ws.send(JSON.stringify(command));
	}
}

export default GameClient;
