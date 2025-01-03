package game

import (
	"fmt"
	"math/rand/v2"
)

type RandomDealer struct {
	deck []Card
	cur  int
}

const MAX_DECK_SIZE = NUM_SUITS * NUM_CARD_VALUES

func NewRandomDealer() *RandomDealer {
	return &RandomDealer{
		deck: shuffleDeck(),
		cur:  0,
	}
}

func (d *RandomDealer) DealCard() (Card, error) {
	if d.cur >= MAX_DECK_SIZE {
		return Card{}, fmt.Errorf("deck is empty")
	}

	defer func() { d.cur = d.cur + 1 }()
	return d.deck[d.cur], nil
}

func shuffleDeck() []Card {
	deck := make([]Card, MAX_DECK_SIZE)
	perm := rand.Perm(MAX_DECK_SIZE)

	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	values := []Rank{Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for i, v := range perm {
		deck[i] = Card{
			Suit: suits[v%NUM_SUITS],
			Rank: values[v/NUM_SUITS],
		}
	}

	return deck
}
