import type { Card } from "~/client/card";

export class PlayCardMove {
	readonly command: string = "playCard";
	card: Card;

	constructor(card: Card) {
		this.card = card;
	}
}

export default PlayCardMove;
