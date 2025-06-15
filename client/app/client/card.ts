export enum Suit {
	Spades   = "Spades",
	Hearts   = "Hearts",
	Diamonds = "Diamonds",
	Clubs    = "Clubs"
}

export enum Rank {
	Seven = "7",
	Eight = "8",
	Nine  = "9",
	Ten   = "10",
	Jack  = "J",
	Queen = "Q",
	King  = "K",
	Ace   = "A"
}

export interface Card {
	suit: Suit,
	rank: Rank
}
