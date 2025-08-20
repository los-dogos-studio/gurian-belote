package game

import "fmt"

type Trick struct {
	StartingPlayer PlayerId
	Cards          map[PlayerId]Card
	Trump          Suit
}

type TrickResult struct {
	WinnerPlayer PlayerId
	Points       int
}

var (
	ErrCardNotOwned                = fmt.Errorf("player does not have this card")
	ErrMustPlayLeadSuitCard        = fmt.Errorf("player must play a card of the lead suit")
	ErrMustPlayTrumpCard           = fmt.Errorf("player must play a trump card")
	ErrMustPlayHigherRankTrumpCard = fmt.Errorf("player must play a higher rank trump card")
)

func NewTrick(startingPlayer PlayerId, trump Suit) *Trick {
	return &Trick{
		StartingPlayer: startingPlayer,
		Cards:          make(map[PlayerId]Card, 4),
		Trump:          trump,
	}
}

func (t *Trick) PlayCard(player PlayerId, card Card, playerCards map[Card]bool) error {
	currentPlayer, err := t.GetCurrentTurn()
	if err != nil {
		return err
	}

	if currentPlayer != player {
		return fmt.Errorf("it is not player's turn")
	}

	if err := t.validateCard(card, playerCards); err != nil {
		return err
	}

	t.Cards[player] = card
	delete(playerCards, card)
	return nil
}

func (t *Trick) GetTrickResult() (*TrickResult, error) {
	if !t.IsFinished() {
		return nil, fmt.Errorf("trick is not finished")
	}

	total := 0
	bestCardOwner := t.StartingPlayer

	for player, card := range t.Cards {
		if card.Suit == t.Trump {
			total += card.Rank.GetTrumpPoints()
		} else {
			total += card.Rank.GetNonTrumpPoints()
		}

		bestCard := t.Cards[bestCardOwner]
		if bestCard.Suit == t.Trump {
			if card.Suit == t.Trump && card.Rank.getTrumpRankOrderIndex() > bestCard.Rank.getTrumpRankOrderIndex() {
				bestCardOwner = player
			}
		} else if card.Suit == t.Trump || (card.Suit == bestCard.Suit && card.Rank.getNonTrumpRankOrderIndex() > bestCard.Rank.getNonTrumpRankOrderIndex()) {
			bestCardOwner = player
		}
	}

	return &TrickResult{bestCardOwner, total}, nil
}

func (t *Trick) IsFinished() bool {
	for _, playerId := range []PlayerId{Player1, Player2, Player3, Player4} {
		if _, ok := t.Cards[playerId]; !ok {
			return false
		}
	}
	return true
}

func (t *Trick) GetCurrentTurn() (PlayerId, error) {
	if t.IsFinished() {
		return Player1, fmt.Errorf("trick is finished")
	}

	return (t.StartingPlayer-Player1+PlayerId(len(t.Cards)))%NUM_PLAYERS + Player1, nil
}

func (t *Trick) GetTableCards() map[PlayerId]Card {
	return t.Cards
}

func (t *Trick) validateCard(card Card, playerCards map[Card]bool) error {
	if owned, ok := playerCards[card]; !ok || !owned {
		return ErrCardNotOwned
	}

	if len(t.Cards) == 0 {
		return nil
	}

	leadSuit := t.getLeadSuit()
	var requiredSuit Suit

	if hasCardOfSuit(playerCards, leadSuit) {
		requiredSuit = leadSuit
	} else if hasCardOfSuit(playerCards, t.Trump) {
		requiredSuit = t.Trump
	} else {
		return nil
	}

	if card.Suit != requiredSuit {
		if requiredSuit == leadSuit {
			return ErrMustPlayLeadSuitCard
		}
		return ErrMustPlayTrumpCard
	}

	if requiredSuit == t.Trump {
		if err := t.validateHigherTrumpRule(card, playerCards); err != nil {
			return err
		}
	}

	return nil
}

func (t *Trick) validateHigherTrumpRule(card Card, playerCards map[Card]bool) error {
	playersHighestTrump := getPlayersHighestTrump(playerCards, t.Trump)

	if playersHighestTrump == nil || card.Suit != t.Trump {
		return nil
	}

	highestTrumpInTrick := t.getHighestTrumpInTrick()

	if highestTrumpInTrick == nil {
		return nil
	}

	if playersHighestTrump.getTrumpRankOrderIndex() > highestTrumpInTrick.getTrumpRankOrderIndex() &&
		card.Rank.getTrumpRankOrderIndex() < highestTrumpInTrick.getTrumpRankOrderIndex() {
		return ErrMustPlayHigherRankTrumpCard
	}

	return nil
}

func hasCardOfSuit(playerCards map[Card]bool, suit Suit) bool {
	for card, owned := range playerCards {
		if owned && card.Suit == suit {
			return true
		}
	}
	return false
}

func (t *Trick) getLeadSuit() Suit {
	return t.Cards[t.StartingPlayer].Suit
}

func (t *Trick) getHighestTrumpInTrick() *Rank {
	var highestRank *Rank = nil

	for _, card := range t.Cards {
		if card.Suit != t.Trump {
			continue
		}

		if highestRank == nil || card.Rank.getTrumpRankOrderIndex() > highestRank.getTrumpRankOrderIndex() {
			highestRank = &card.Rank
		}
	}

	return highestRank
}

func getPlayersHighestTrump(playerCards map[Card]bool, trump Suit) *Rank {
	var highestRank *Rank = nil

	for card, owned := range playerCards {
		if !owned || card.Suit != trump {
			continue
		}

		if highestRank == nil || card.Rank.getTrumpRankOrderIndex() > highestRank.getTrumpRankOrderIndex() {
			highestRank = &card.Rank
		}
	}

	return highestRank
}
