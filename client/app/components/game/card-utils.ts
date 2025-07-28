import { Suit } from "~/client/card";

export function getSuitSymbol(suit: Suit): "♠" | "♥" | "♦" | "♣" {
	switch (suit) {
		case Suit.Spades:
			return "♠";
		case Suit.Hearts:
			return "♥";
		case Suit.Diamonds:
			return "♦";
		case Suit.Clubs:
			return "♣";
		default:
			throw new Error("Invalid suit");
	}
}

export function getSuitColor(suit: Suit): string {
	switch (suit) {
		case Suit.Spades:
			return "text-black";
		case Suit.Hearts:
			return "text-red-500";
		case Suit.Diamonds:
			return "text-red-500";
		case Suit.Clubs:
			return "text-black";
		default:
			throw new Error("Invalid suit");
	}
}

