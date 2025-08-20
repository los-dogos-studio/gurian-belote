package game

import (
	"testing"
)

func TestValidateCard(t *testing.T) {
	testCases := []struct {
		name          string
		trick         *Trick
		playerCards   map[Card]bool
		cardToPlay    Card
		expectedError error
	}{
		{
			name: "Player does not own the card",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards:          make(map[PlayerId]Card),
				Trump:          Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Spades, Rank: Ace}: true,
			},
			cardToPlay:    Card{Suit: Hearts, Rank: Ace},
			expectedError: ErrCardNotOwned,
		},
		{
			name: "First card of the trick",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards:          make(map[PlayerId]Card),
				Trump:          Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Clubs, Rank: Ace}: true,
			},
			cardToPlay:    Card{Suit: Clubs, Rank: Ace},
			expectedError: nil,
		},
		{
			name: "Must follow lead suit - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Clubs, Rank: Seven},
					Player2: {Suit: Diamonds, Rank: Nine},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Clubs, Rank: Queen}:   true,
				{Suit: Diamonds, Rank: Jack}: true,
			},
			cardToPlay:    Card{Suit: Clubs, Rank: Queen},
			expectedError: nil,
		},
		{
			name: "Must follow lead suit - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Hearts, Rank: Seven},
					Player2: {Suit: Diamonds, Rank: Eight},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Hearts, Rank: Ace}:   true,
				{Suit: Diamonds, Rank: Ten}: true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Ten},
			expectedError: ErrMustPlayLeadSuitCard,
		},
		{
			name: "Cannot follow lead suit, must play trump - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Hearts, Rank: Seven},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Clubs, Rank: Eight}:    true,
				{Suit: Diamonds, Rank: Eight}: true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Eight},
			expectedError: nil,
		},
		{
			name: "Cannot follow lead suit, must play trump - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Hearts, Rank: Seven},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Clubs, Rank: Eight}:    true,
				{Suit: Diamonds, Rank: Eight}: true,
			},
			cardToPlay:    Card{Suit: Clubs, Rank: Eight},
			expectedError: ErrMustPlayTrumpCard,
		},
		{
			name: "Cannot follow lead suit or trump, can play anything",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Clubs, Rank: Ten},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Hearts, Rank: Nine}: true,
				{Suit: Spades, Rank: Jack}: true,
			},
			cardToPlay:    Card{Suit: Hearts, Rank: Nine},
			expectedError: nil,
		},
		{
			name: "Must play higher trump - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Hearts, Rank: Eight},
					Player2: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: King}:  true,
				{Suit: Diamonds, Rank: Eight}: true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: King},
			expectedError: nil,
		},
		{
			name: "Must play higher trump - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Hearts, Rank: Eight},
					Player2: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: King}:  true,
				{Suit: Diamonds, Rank: Eight}: true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Eight},
			expectedError: ErrMustPlayHigherRankTrumpCard,
		},
		{
			name: "Does not have higher trump, must play lower trump - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Spades, Rank: Ace},
					Player2: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Clubs, Rank: Ten}:      true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Eight},
			expectedError: nil,
		},
		{
			name: "Does not have higher trump, must play lower trump - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Clubs, Rank: Ace},
					Player2: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Spades, Rank: Ten}:     true,
			},
			cardToPlay:    Card{Suit: Spades, Rank: Ten},
			expectedError: ErrMustPlayTrumpCard,
		},
		{
			name: "Lead card is trump, must play higher trump - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Diamonds, Rank: King}:  true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: King},
			expectedError: nil,
		},
		{
			name: "Lead card is trump, must play higher trump - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Diamonds, Rank: King}:  true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Eight},
			expectedError: ErrMustPlayHigherRankTrumpCard,
		},
		{
			name: "Lead card is trump, does not have higher trump, must play lower trump - success",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Clubs, Rank: Eight}:    true,
			},
			cardToPlay:    Card{Suit: Diamonds, Rank: Eight},
			expectedError: nil,
		},
		{
			name: "Lead card is trump, does not have higher trump, must play lower trump - error",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Diamonds, Rank: Eight}: true,
				{Suit: Clubs, Rank: Eight}:    true,
			},
			cardToPlay:    Card{Suit: Clubs, Rank: Eight},
			expectedError: ErrMustPlayLeadSuitCard,
		},
		{
			name: "Lead card is trump, player has no trump, can play any card",
			trick: &Trick{
				StartingPlayer: Player1,
				Cards: map[PlayerId]Card{
					Player1: {Suit: Diamonds, Rank: Queen},
				},
				Trump: Diamonds,
			},
			playerCards: map[Card]bool{
				{Suit: Hearts, Rank: Ace}: true,
			},
			cardToPlay:    Card{Suit: Hearts, Rank: Ace},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.trick.validateCard(tc.cardToPlay, tc.playerCards)
			if err != tc.expectedError {
				t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
