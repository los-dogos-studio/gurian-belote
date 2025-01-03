package game

import "fmt"

type Hand struct {
	state          HandState
	currentTrick   *Trick
	startingPlayer PlayerId
	totals         map[TeamId]int
	playerCards    map[PlayerId]map[Card]bool
	dealer         Dealer

	tableTrumpCard            Card
	tableTrumpSelectionStatus map[PlayerId]bool
	freeTrumpSelectionStatus  map[PlayerId]bool

	trump Suit
}

type Dealer interface {
	DealCard() (Card, error)
}

type HandState string

const (
	TableTrumpSelection HandState = "TableTrumpSelection"
	FreeTrumpSelection  HandState = "FreeTrumpSelection"
	HandInProgress      HandState = "HandInProgress"
	HandFinished        HandState = "HandFinished"
)

func (h *Hand) GetTrick() *Trick {
	if h.state == HandInProgress {
		// TODO: copy?
		return h.currentTrick
	}
	return nil
}

func NewHand(startingPlayer PlayerId, dealer Dealer) *Hand {
	hand := &Hand{
		state:                     TableTrumpSelection,
		currentTrick:              nil,
		startingPlayer:            startingPlayer,
		totals:                    nil,
		playerCards:               makePlayerCards(),
		dealer:                    dealer,
		tableTrumpCard:            Card{},
		tableTrumpSelectionStatus: map[PlayerId]bool{},
		freeTrumpSelectionStatus:  map[PlayerId]bool{},
		trump:                     Spades,
	}

	hand.totals = map[TeamId]int{}
	hand.totals[Team1] = 0
	hand.totals[Team2] = 0

	tableTrumpCard, err := dealer.DealCard()
	if err != nil {
		panic(err)
	}

	hand.tableTrumpCard = tableTrumpCard

	return hand
}

func (h *Hand) PlayCard(player PlayerId, card Card) error {
	if h.state != HandInProgress {
		return fmt.Errorf("hand is not in progress")
	}

	if err := h.currentTrick.PlayCard(player, card, h.playerCards[player]); err != nil {
		return err
	}

	if h.currentTrick.IsFinished() {
		trickResult, err := h.currentTrick.GetTrickResult()
		if err != nil {
			panic(err)
		}
		h.handleTrickResult(trickResult)
	}

	return nil
}

func (h *Hand) AcceptTableTrump(player PlayerId, accept bool) error {
	if h.state != TableTrumpSelection {
		return fmt.Errorf("table trump selection is not in progress, current state: %s", h.state)
	}

	if err := h.checkIsTrumpSelectionTurnFor(player, h.tableTrumpSelectionStatus); err != nil {
		return err
	}

	if accept {
		h.playerCards[player][h.tableTrumpCard] = true
		h.handleTrumpSelected(h.tableTrumpCard.Suit)
		return nil
	}

	h.tableTrumpSelectionStatus[player] = true
	if len(h.tableTrumpSelectionStatus) == NUM_PLAYERS {
		h.state = FreeTrumpSelection
	}

	return nil
}

func (h *Hand) SelectTrump(player PlayerId, suit *Suit) error {
	if h.state != FreeTrumpSelection {
		return fmt.Errorf("free trump selection is not in progress")
	}

	if err := h.checkIsTrumpSelectionTurnFor(player, h.freeTrumpSelectionStatus); err != nil {
		return err
	}

	if suit == &h.tableTrumpCard.Suit {
		return fmt.Errorf("trump suit cannot be the same as table trump suit")
	}

	if player == h.getLastPlayer() && suit == nil {
		return fmt.Errorf("final player must select a trump suit")
	}

	if suit != nil {
		h.playerCards[player][h.tableTrumpCard] = true
		h.handleTrumpSelected(*suit)
		return nil
	}

	h.freeTrumpSelectionStatus[player] = true
	return nil
}

// TODO: copy?
// TODO: hide?
func (h *Hand) GetPlayerCards(player PlayerId) map[Card]bool {
	return h.playerCards[player]
}

func (h *Hand) handleTrickResult(trickResult *TrickResult) {
	h.totals[trickResult.winnerPlayer.GetTeam()] += trickResult.points

	if h.checkEndCondition() {
		h.state = HandFinished
		return
	}

	h.currentTrick = NewTrick(trickResult.winnerPlayer, h.trump)
}

func (h *Hand) checkEndCondition() bool {
	return len(h.playerCards[Player1]) == 0
}

func (h *Hand) checkIsTrumpSelectionTurnFor(player PlayerId, selections map[PlayerId]bool) error {
	if selections[player] {
		return fmt.Errorf("player has already selected")
	}

	if h.startingPlayer == player {
		return nil
	}

	fmt.Println(h.startingPlayer, player)
	if selected := selections[(player-Player1+NUM_PLAYERS-1)%NUM_PLAYERS+Player1]; !selected {
		fmt.Println((player-Player1+NUM_PLAYERS-1)%NUM_PLAYERS + Player1)
		return fmt.Errorf("previous player has not selected yet")
	}

	return nil
}

func (h *Hand) handleTrumpSelected(trump Suit) {
	h.trump = trump
	h.state = HandInProgress
	h.currentTrick = NewTrick(h.startingPlayer, h.trump)
	h.dealCards()
}

func (h *Hand) dealCards() {
	for player := Player1; player <= Player4; player++ {
		fmt.Println("Dealing cards for player", player)
		for len(h.playerCards[player]) < NUM_CARD_PER_PLAYER {
			cardDealt, err := h.dealer.DealCard()
			if err != nil {
				panic(err)
			}
			h.playerCards[player][cardDealt] = true
		}
		fmt.Println("Player", player, "has cards", h.playerCards[player])
	}
}

func makePlayerCards() map[PlayerId]map[Card]bool {
	playerCards := make(map[PlayerId]map[Card]bool, NUM_PLAYERS)
	for player := Player1; player <= Player4; player++ {
		playerCards[player] = make(map[Card]bool, NUM_CARD_PER_PLAYER)
	}
	return playerCards
}

func (h *Hand) getLastPlayer() PlayerId {
	return Player1 + (h.startingPlayer-Player1+NUM_PLAYERS-1)%NUM_PLAYERS
}
