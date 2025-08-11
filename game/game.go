package game

import "fmt"

// NOT thread-safe
type BeloteGame struct {
	state          GameState
	scores         map[TeamId]int
	startingPlayer PlayerId
	targetScore    int

	currentHand *Hand
	handNumber  int
}

const (
	NUM_PLAYERS            = 4
	NUM_CARDS_PER_PLAYER   = 8
	NUM_CARDS_BEFORE_TRUMP = 5
	TARGET_SCORE           = 1000
)

type GameState string

const (
	GameReady      GameState = "Ready"
	GameInProgress GameState = "InProgress"
	GameFinished   GameState = "Finished"
)

type PlayerId int

const (
	NoPlayerId PlayerId = iota
	Player1
	Player2
	Player3
	Player4
)

type TeamId int

const (
	NoTeamId TeamId = iota
	Team1
	Team2
)

func (p PlayerId) GetTeam() TeamId {
	if p == Player1 || p == Player3 {
		return Team1
	}
	return Team2
}

func (p PlayerId) GetTeammateId() PlayerId {
	if p == Player1 {
		return Player3
	} else if p == Player3 {
		return Player1
	} else if p == Player2 {
		return Player4
	}
	return Player2
}

func (p PlayerId) GetNextPlayerId() PlayerId {
	return (p-1+1)%4 + 1
}

func (p PlayerId) GetPreviousPlayerId() PlayerId {
	return (p-1+3)%4 + 1
}

func NewBeloteGame() BeloteGame {
	scores := make(map[TeamId]int)
	scores[Team1] = 0
	scores[Team2] = 0

	// TODO: hand factory?
	return BeloteGame{
		state:          GameReady,
		scores:         scores,
		startingPlayer: Player1,
		targetScore:    TARGET_SCORE,
		currentHand:    nil,
		handNumber:     0,
	}
}

func (gm *BeloteGame) Start() {
	gm.state = GameInProgress
	gm.setupHand()
}

func (gm *BeloteGame) PlayCard(player PlayerId, card Card) error {
	if gm.state != GameInProgress {
		return fmt.Errorf("game is not in progress")
	}

	err := gm.currentHand.PlayCard(player, card)
	if err != nil {
		return err
	}

	if gm.currentHand.State == HandFinished {
		gm.handleHandEnd()
	}

	return nil
}

func (gm *BeloteGame) AcceptTableTrump(player PlayerId, accept bool) error {
	if gm.state != GameInProgress {
		return fmt.Errorf("game is not in progress")
	}

	return gm.currentHand.AcceptTableTrump(player, accept)
}

func (gm *BeloteGame) SelectTrump(player PlayerId, suit *Suit) error {
	if gm.state != GameInProgress {
		return fmt.Errorf("game is not in progress")
	}

	return gm.currentHand.SelectTrump(player, suit)
}

// TODO: return a copy?
// TODO: decide exported functions
func (gm *BeloteGame) GetState() GameState {
	return gm.state
}

func (gm *BeloteGame) GetHand() *Hand {
	return gm.currentHand
}

func (gm *BeloteGame) GetScores() map[TeamId]int {
	return gm.scores
}

func (gm *BeloteGame) setupHand() {
	gm.currentHand = NewHand(calculateHandStartingPlayer(gm.startingPlayer, gm.handNumber), NewRandomDealer())
}

func calculateHandStartingPlayer(startingPlayer PlayerId, handNumber int) PlayerId {
	return Player1 + (startingPlayer-Player1+PlayerId(handNumber))%NUM_PLAYERS
}

func (gm *BeloteGame) handleHandEnd() {
	gm.scores[Team1] += gm.currentHand.Totals[Team1]
	gm.scores[Team2] += gm.currentHand.Totals[Team2]

	if gm.checkEndCondition() {
		gm.state = GameFinished
		gm.currentHand = nil
		return
	}

	gm.refreshHand()
}

func (gm *BeloteGame) refreshHand() {
	gm.handNumber++
	gm.setupHand()
}

func (gm *BeloteGame) checkEndCondition() bool {
	return gm.scores[Team1] >= gm.targetScore || gm.scores[Team2] >= gm.targetScore
}
