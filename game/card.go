package game

type Suit string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

const NUM_SUITS = 4

type Rank string

const (
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
	Ace   Rank = "A"
)

const NUM_CARD_VALUES = 8

type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
}

func (c *Card) String() string {
	return string(c.Rank) + " of " + string(c.Suit)
}

var naturalOrderIndex = map[Rank]int{
	Seven: 0,
	Eight: 1,
	Nine:  2,
	Ten:   3,
	Jack:  4,
	Queen: 5,
	King:  6,
	Ace:   7,
}

var trumpTrickOrderIndex = map[Rank]int{
	Seven: 0,
	Eight: 1,
	Queen: 2,
	King:  3,
	Ten:   4,
	Ace:   5,
	Nine:  6,
	Jack:  7,
}

var nonTrumpTrickOrderIndex = map[Rank]int{
	Seven: 0,
	Eight: 1,
	Nine:  2,
	Jack:  3,
	Queen: 4,
	King:  5,
	Ten:   6,
	Ace:   7,
}

func (r *Rank) NaturalOrder() int {
	return naturalOrderIndex[*r]
}

func (r *Rank) TrickOrder(isTrump bool) int {
	if isTrump {
		return trumpTrickOrderIndex[*r]
	}
	return nonTrumpTrickOrderIndex[*r]
}

func (r *Rank) GetNonTrumpPoints() int {
	points := map[Rank]int{
		Seven: 0,
		Eight: 0,
		Nine:  0,
		Jack:  2,
		Queen: 3,
		King:  4,
		Ten:   10,
		Ace:   11,
	}
	return points[*r]
}

func (r *Rank) GetTrumpPoints() int {
	points := map[Rank]int{
		Seven: 0,
		Eight: 0,
		Queen: 3,
		King:  4,
		Ten:   10,
		Ace:   11,
		Nine:  14,
		Jack:  20,
	}
	return points[*r]
}

func Less(r1, r2 Rank, isTrump bool) bool {
	return r1.TrickOrder(isTrump) < r2.TrickOrder(isTrump)
}
