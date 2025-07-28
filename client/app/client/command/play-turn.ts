import type { GameMove } from "./move/game-move";

export class PlayTurnCommand {
	readonly command: string = "playTurn";
	move: GameMove;

	constructor(move: GameMove) {
		this.move = move;
	}
}

export default PlayTurnCommand;
