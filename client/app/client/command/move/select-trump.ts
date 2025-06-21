import type { Suit } from "~/client/card";

export class SelectTrumpMove {
	readonly command: string = "selectTrump";
	suit: Suit | null;

	constructor(suit: Suit | null) {
		this.suit = suit;
	}
}

export default SelectTrumpMove;
