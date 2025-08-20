import { IsEnum } from "class-validator";

export enum Suit {
	Spades = "Spades",
	Hearts = "Hearts",
	Diamonds = "Diamonds",
	Clubs = "Clubs"
}

export enum Rank {
	Seven = "7",
	Eight = "8",
	Nine = "9",
	Ten = "10",
	Jack = "J",
	Queen = "Q",
	King = "K",
	Ace = "A"
}

const trumpOrder = [Rank.Jack, Rank.Nine, Rank.Ace, Rank.Ten, Rank.King, Rank.Queen, Rank.Eight, Rank.Seven];
const nonTrumpOrder = [Rank.Ace, Rank.Ten, Rank.King, Rank.Queen, Rank.Jack, Rank.Nine, Rank.Eight, Rank.Seven];

export class Card {
	@IsEnum(Suit)
	suit: Suit;

	@IsEnum(Rank)
	rank: Rank;

	constructor(suit: Suit, rank: Rank) {
		this.suit = suit;
		this.rank = rank;
	}

	compare(other: Card, trumpSuit: Suit): number {
		const thisOrder = this.suit === trumpSuit ? trumpOrder : nonTrumpOrder;
		const otherOrder = other.suit === trumpSuit ? trumpOrder : nonTrumpOrder;

		const thisRankIndex = thisOrder.indexOf(this.rank);
		const otherRankIndex = otherOrder.indexOf(other.rank);

		if (this.suit === other.suit) {
			return otherRankIndex - thisRankIndex;
		}

		if (this.suit === trumpSuit) {
			return 1;
		}

		if (other.suit === trumpSuit) {
			return -1;
		}

		return 0;
	}

	equals(other: Card) {
		return this.suit === other.suit && this.rank === other.rank;
	}
}
