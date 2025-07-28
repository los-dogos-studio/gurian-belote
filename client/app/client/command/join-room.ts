class JoinRoomCommand {
	readonly command: string = 'joinRoom';
	roomId: string;

	constructor(roomId: string) {
		this.roomId = roomId;
	}
}

export default JoinRoomCommand;
