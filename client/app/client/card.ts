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

export class Card {
	@IsEnum(Suit)
	suit: Suit;

	@IsEnum(Rank)
	rank: Rank;

	constructor(suit: Suit, rank: Rank) {
		this.suit = suit;
		this.rank = rank;
	}
}
