package game

import "fmt"

type Hand struct {
	State          HandState
	CurrentTrick   *Trick
	StartingPlayer PlayerId
	Totals         map[TeamId]int
	PlayerCards    map[PlayerId]map[Card]bool

	TableTrumpCard            Card
	TableTrumpSelectionStatus map[PlayerId]bool
	FreeTrumpSelectionStatus  map[PlayerId]bool

	Trump Suit

	dealer Dealer
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

func NewHand(startingPlayer PlayerId, dealer Dealer) *Hand {
	hand := &Hand{
		State:                     TableTrumpSelection,
		CurrentTrick:              nil,
		StartingPlayer:            startingPlayer,
		Totals:                    nil,
		PlayerCards:               makePlayerCards(),
		dealer:                    dealer,
		TableTrumpCard:            Card{},
		TableTrumpSelectionStatus: map[PlayerId]bool{},
		FreeTrumpSelectionStatus:  map[PlayerId]bool{},
		Trump:                     Spades,
	}

	hand.Totals = map[TeamId]int{}
	hand.Totals[Team1] = 0
	hand.Totals[Team2] = 0

	tableTrumpCard, err := dealer.DealCard()
	if err != nil {
		panic(err)
	}

	if tableTrumpCard.Rank == Jack {
		hand.handleTrumpSelected(tableTrumpCard.Suit)
		return hand
	}

	hand.TableTrumpCard = tableTrumpCard
	hand.dealInitialCards()

	return hand
}

func (h *Hand) PlayCard(player PlayerId, card Card) error {
	if h.State != HandInProgress {
		return fmt.Errorf("hand is not in progress")
	}

	if err := h.CurrentTrick.PlayCard(player, card, h.PlayerCards[player]); err != nil {
		return err
	}

	if h.CurrentTrick.IsFinished() {
		trickResult, err := h.CurrentTrick.GetTrickResult()
		if err != nil {
			panic(err)
		}
		h.handleTrickResult(trickResult)
	}

	return nil
}

func (h *Hand) AcceptTableTrump(player PlayerId, accept bool) error {
	if h.State != TableTrumpSelection {
		return fmt.Errorf("table trump selection is not in progress, current state: %s", h.State)
	}

	if err := h.checkIsTrumpSelectionTurnFor(player, h.TableTrumpSelectionStatus); err != nil {
		return err
	}

	if accept {
		h.PlayerCards[player][h.TableTrumpCard] = true
		h.handleTrumpSelected(h.TableTrumpCard.Suit)
		return nil
	}

	h.TableTrumpSelectionStatus[player] = true
	if len(h.TableTrumpSelectionStatus) == NUM_PLAYERS {
		h.State = FreeTrumpSelection
	}

	return nil
}

func (h *Hand) SelectTrump(player PlayerId, suit *Suit) error {
	if h.State != FreeTrumpSelection {
		return fmt.Errorf("free trump selection is not in progress")
	}

	if err := h.checkIsTrumpSelectionTurnFor(player, h.FreeTrumpSelectionStatus); err != nil {
		return err
	}

	if player == h.getLastPlayer() && suit == nil {
		return fmt.Errorf("final player must select a trump suit")
	}

	if suit == nil {
		h.FreeTrumpSelectionStatus[player] = true
		return nil
	}

	if *suit == h.TableTrumpCard.Suit {
		return fmt.Errorf("trump suit cannot be the same as table trump suit")
	}

	h.PlayerCards[player][h.TableTrumpCard] = true
	h.handleTrumpSelected(*suit)
	return nil
}

func (h *Hand) GetTrick() *Trick {
	if h.State == HandInProgress {
		// TODO: copy?
		return h.CurrentTrick
	}
	return nil
}

// TODO: copy?
// TODO: hide?
func (h *Hand) GetPlayerCards(player PlayerId) map[Card]bool {
	return h.PlayerCards[player]
}

func (h *Hand) GetTotals() map[TeamId]int {
	return h.Totals
}

func (h *Hand) GetState() HandState {
	return h.State
}

func (h *Hand) GetTrump() Suit {
	return h.Trump
}

func (h *Hand) GetTableTrump() Card {
	return h.TableTrumpCard
}

func (h *Hand) GetCurrentTurn() (PlayerId, error) {
	switch h.State {
	case TableTrumpSelection:
		return h.getCurrentTrumpSelectionTurn(h.TableTrumpSelectionStatus)
	case FreeTrumpSelection:
		return h.getCurrentTrumpSelectionTurn(h.FreeTrumpSelectionStatus)
	case HandInProgress:
		return h.CurrentTrick.GetCurrentTurn()
	case HandFinished:
		return Player1, fmt.Errorf("hand is finished")
	default:
		panic(fmt.Sprintf("unexpected game.HandState: %#v", h.State))
	}
}

func (h *Hand) dealInitialCards() {
	for player := Player1; player <= Player4; player += 1 {
		for len(h.PlayerCards[player]) < NUM_CARDS_BEFORE_TRUMP {
			cardDealt, err := h.dealer.DealCard()
			if err != nil {
				panic(err)
			}
			h.PlayerCards[player][cardDealt] = true
		}
	}
}

func (h *Hand) handleTrickResult(trickResult *TrickResult) {
	h.Totals[trickResult.WinnerPlayer.GetTeam()] += trickResult.Points

	if h.checkEndCondition() {
		h.State = HandFinished
		return
	}

	h.CurrentTrick = NewTrick(trickResult.WinnerPlayer, h.Trump)
}

func (h *Hand) checkEndCondition() bool {
	return len(h.PlayerCards[Player1]) == 0
}

func (h *Hand) getCurrentTrumpSelectionTurn(selections map[PlayerId]bool) (PlayerId, error) {
	for i := 0; i < NUM_PLAYERS; i++ {
		player := Player1 + (h.StartingPlayer-Player1+PlayerId(i))%NUM_PLAYERS
		if !selections[player] {
			return player, nil
		}
	}
	return Player1, fmt.Errorf("all players have selected")
}

func (h *Hand) checkIsTrumpSelectionTurnFor(player PlayerId, selections map[PlayerId]bool) error {
	if selections[player] {
		return fmt.Errorf("player has already selected")
	}

	currentPlayer, err := h.getCurrentTrumpSelectionTurn(selections)
	if err != nil {
		panic(err)
	}

	if player != currentPlayer {
		return fmt.Errorf("not player's turn")
	}

	return nil
}

func (h *Hand) handleTrumpSelected(trump Suit) {
	h.Trump = trump
	h.State = HandInProgress
	h.CurrentTrick = NewTrick(h.StartingPlayer, h.Trump)
	h.dealCards()
}

func (h *Hand) dealCards() {
	for player := Player1; player <= Player4; player++ {
		for len(h.PlayerCards[player]) < NUM_CARDS_PER_PLAYER {
			cardDealt, err := h.dealer.DealCard()
			if err != nil {
				panic(err)
			}
			h.PlayerCards[player][cardDealt] = true
		}
	}
}

func makePlayerCards() map[PlayerId]map[Card]bool {
	playerCards := make(map[PlayerId]map[Card]bool, NUM_PLAYERS)
	for player := Player1; player <= Player4; player++ {
		playerCards[player] = make(map[Card]bool, NUM_CARDS_PER_PLAYER)
	}
	return playerCards
}

func (h *Hand) getLastPlayer() PlayerId {
	return Player1 + (h.StartingPlayer-Player1+NUM_PLAYERS-1)%NUM_PLAYERS
}
