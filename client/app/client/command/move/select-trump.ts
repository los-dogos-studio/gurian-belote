import type { Rank } from "~/client/card";

export class SelectTrumpMove {
	readonly command: string = "selectTrump";
	suit: Rank | null;

	constructor(suit: Rank | null) {
		this.suit = suit;
	}
}

export default SelectTrumpMove;
