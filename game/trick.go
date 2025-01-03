package game

import "fmt"

type Trick struct {
	startingPlayer PlayerId
	cards          map[PlayerId]Card
	trump          Suit
}

type TrickResult struct {
	winnerPlayer PlayerId
	points       int
}

func NewTrick(startingPlayer PlayerId, trump Suit) *Trick {
	return &Trick{
		startingPlayer: startingPlayer,
		cards:          make(map[PlayerId]Card, 4),
		trump:          trump,
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

	t.cards[player] = card
	delete(playerCards, card)
	return nil
}

func (t *Trick) GetTrickResult() (*TrickResult, error) {
	if !t.IsFinished() {
		return nil, fmt.Errorf("trick is not finished")
	}

	total := 0
	bestCardOwner := t.startingPlayer

	for player, card := range t.cards {
		if card.Suit == t.trump {
			total += card.Rank.GetTrumpPoints()
		} else {
			total += card.Rank.GetNonTrumpPoints()
		}

		bestCard := t.cards[bestCardOwner]
		if bestCard.Suit == t.trump {
			if card.Suit == t.trump && card.Rank.getTrumpRankOrderIndex() > bestCard.Rank.getTrumpRankOrderIndex() {
				bestCardOwner = player
			}
		} else if (card.Suit == t.trump || (card.Suit == bestCard.Suit && card.Rank.getNonTrumpRankOrderIndex() > bestCard.Rank.getNonTrumpRankOrderIndex())) {
			bestCardOwner = player
		}
	}

	return &TrickResult{bestCardOwner, total}, nil
}

func (t *Trick) IsFinished() bool {
	for _, playerId := range []PlayerId{Player1, Player2, Player3, Player4} {
		if _, ok := t.cards[playerId]; !ok {
			return false
		}
	}
	return true
}

func (t *Trick) GetCurrentTurn() (PlayerId, error) {
	if t.IsFinished() {
		return Player1, fmt.Errorf("trick is finished")
	}

	return (t.startingPlayer-Player1+PlayerId(len(t.cards)))%NUM_PLAYERS + Player1, nil
}

func (t *Trick) validateCard(card Card, playerCards map[Card]bool) error {
	if owned, ok := playerCards[card]; !ok || !owned {
		return fmt.Errorf("player does not have this card")
	}

	if len(t.cards) == 0 {
		return nil
	}

	highestTrumpInPlay := t.getHighestTrumpInPlay()

	if highestTrumpInPlay != nil {
		if err := t.validateHigherTrumpRule(card, playerCards); err != nil {
			return err
		}
	}

	return t.validateIsSuitFollowed(card, playerCards)
}

func (t *Trick) validateIsSuitFollowed(card Card, playerCards map[Card]bool) error {
	originalSuit := t.getOriginalSuit()
	if card.Suit == originalSuit {
		return nil
	}
	if hasCardOfSuit(playerCards, originalSuit) {
		return fmt.Errorf("player must play a card of the original suit")
	}

	if card.Suit == t.trump {
		return nil
	}

	if hasCardOfSuit(playerCards, t.trump) {
		return fmt.Errorf("player must play a trump card")
	}
	return nil
}

func (t *Trick) validateHigherTrumpRule(card Card, playerCards map[Card]bool) error {
	playersHighestTrump := getPlayersHighestTrump(playerCards, t.trump)
	highestTrumpInPlay := t.getHighestTrumpInPlay()

	if playersHighestTrump.getTrumpRankOrderIndex() > highestTrumpInPlay.getTrumpRankOrderIndex() &&
		(card.Suit != t.trump || card.Rank.getTrumpRankOrderIndex() < playersHighestTrump.getTrumpRankOrderIndex()) {
		return fmt.Errorf("player must play a higher trump card")
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

func (t *Trick) getOriginalSuit() Suit {
	return t.cards[t.startingPlayer].Suit
}

func (t *Trick) getHighestTrumpInPlay() *Rank {
	var highestRank *Rank = nil

	for _, card := range t.cards {
		if card.Suit != t.trump {
			continue
		}

		if highestRank == nil || card.Rank.getTrumpRankOrderIndex() > highestRank.getTrumpRankOrderIndex() {
			highestRank = &card.Rank
		}
	}

	return nil
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
