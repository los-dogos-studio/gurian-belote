package room

import (
	"errors"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type GameState string

type TrickDump struct {
	PlayedCards    map[game.PlayerId]game.Card `json:"playedCards"`
	StartingPlayer game.PlayerId               `json:"startingPlayer"`
}

type HandDump interface {
	GetState() game.HandState
	GetStartingPlayer() game.PlayerId
}

type TableTrumpSelectionHandDump struct {
	State           game.HandState         `json:"state"`
	TableTrumpCard  game.Card              `json:"tableTrumpCard"`
	SelectionStatus map[game.PlayerId]bool `json:"selectionStatus"`
	StartingPlayer  game.PlayerId          `json:"startingPlayer"`
	PreviousTrick   *TrickDump             `json:"previousTrick,omitempty"`
}

func (d *TableTrumpSelectionHandDump) GetState() game.HandState {
	return d.State
}

func (d *TableTrumpSelectionHandDump) GetStartingPlayer() game.PlayerId {
	return d.StartingPlayer
}

type FreeTrumpSelectionHandDump struct {
	State           game.HandState         `json:"state"`
	TableTrumpCard  game.Card              `json:"tableTrumpCard"`
	SelectionStatus map[game.PlayerId]bool `json:"selectionStatus"`
	StartingPlayer  game.PlayerId          `json:"startingPlayer"`
	PreviousTrick   *TrickDump             `json:"previousTrick,omitempty"`
}

func (d *FreeTrumpSelectionHandDump) GetState() game.HandState {
	return d.State
}

func (d *FreeTrumpSelectionHandDump) GetStartingPlayer() game.PlayerId {
	return d.StartingPlayer
}

type InProgressHandDump struct {
	State         game.HandState      `json:"state"`
	Trump         game.Suit           `json:"trump"`
	Trick         TrickDump           `json:"trick"`
	PreviousTrick *TrickDump          `json:"previousTrick,omitempty"`
	Totals        map[game.TeamId]int `json:"totals"`
}

func (d *InProgressHandDump) GetState() game.HandState {
	return d.State
}

func (d *InProgressHandDump) GetStartingPlayer() game.PlayerId {
	return d.Trick.StartingPlayer
}

type StateDump struct {
	RoomId    string                   `json:"roomId"`
	Players   map[game.PlayerId]string `json:"players"`
	Teams     map[game.TeamId][]string `json:"teams"`
	Hand      HandDump                 `json:"hand"`
	GameState game.GameState           `json:"gameState"`
	Scores    map[game.TeamId]int      `json:"scores"`
}

type UserStateDump struct {
	GameState StateDump     `json:"gameState"`
	UserId    string        `json:"userId"`
	PlayerId  game.PlayerId `json:"playerId"`
	UserCards []game.Card   `json:"userCards"`
}

var (
	ErrUserNotInRoom = errors.New("user not in room")
)

func (r *Room) DumpState() StateDump {
	return StateDump{
		RoomId:    r.Id,
		Players:   r.dumpPlayersMap(),
		Teams:     r.dumpTeams(),
		Hand:      r.dumpHand(),
		GameState: r.dumpGameState(),
		Scores:    r.dumpScore(),
	}
}

func (r *Room) DumpUserState(userId string) (UserStateDump, error) {
	_, ok := r.Users[userId]
	if !ok {
		return UserStateDump{}, ErrUserNotInRoom
	}

	state := r.DumpState()
	userCards := r.dumpUserCards(userId)

	return UserStateDump{
		GameState: state,
		UserId:    userId,
		PlayerId:  r.Users[userId].playerId,
		UserCards: userCards,
	}, nil
}

func (r *Room) dumpUserCards(userId string) []game.Card {
	user, ok := r.Users[userId]
	if !ok {
		return nil
	}

	hand := r.Game.GetHand()
	if hand == nil {
		return nil
	}

	cards := hand.GetPlayerCards(user.playerId)
	result := make([]game.Card, 0, len(cards))
	for card, exists := range cards {
		if exists {
			result = append(result, card)
		}
	}

	return result
}

func (r *Room) dumpPlayersMap() map[game.PlayerId]string {
	players := make(map[game.PlayerId]string)
	for id, user := range r.Users {
		players[user.playerId] = id
	}
	return players
}

func (r *Room) dumpTeams() map[game.TeamId][]string {
	if r.Users == nil {
		return nil
	}

	teams := make(map[game.TeamId][]string)

	for id, user := range r.Users {
		teams[user.team] = append(teams[user.team], id)
	}
	return teams
}

func (r *Room) dumpHand() HandDump {
	hand := r.Game.GetHand()
	if hand == nil {
		return nil
	}

	switch hand.GetState() {
	case game.TableTrumpSelection:
		return dumpTableTrumpSelectionHand(hand)
	case game.FreeTrumpSelection:
		return dumpFreeTrumpSelectionHand(hand)
	case game.HandInProgress:
		return dumpInProgressHand(hand)
	case game.HandFinished:
		return nil
	}
	return nil
}

func (r *Room) dumpGameState() game.GameState {
	return r.Game.GetState()
}

func (r *Room) dumpHandState() game.HandState {
	if r.Game.GetHand() == nil {
		return game.HandFinished
	}
	return r.Game.GetHand().GetState()
}

func (r *Room) dumpScore() map[game.TeamId]int {
	scores := make(map[game.TeamId]int)
	for team, score := range r.Game.GetScores() {
		scores[team] = score
	}
	return scores
}

func dumpTableTrumpSelectionHand(hand *game.Hand) *TableTrumpSelectionHandDump {
	return &TableTrumpSelectionHandDump{
		State:           hand.GetState(),
		TableTrumpCard:  hand.TableTrumpCard,
		SelectionStatus: hand.TableTrumpSelectionStatus,
		StartingPlayer:  hand.StartingPlayer,
		PreviousTrick:   dumpTrick(hand.PreviousTrick),
	}
}

func dumpFreeTrumpSelectionHand(hand *game.Hand) *FreeTrumpSelectionHandDump {
	return &FreeTrumpSelectionHandDump{
		State:           hand.GetState(),
		TableTrumpCard:  hand.TableTrumpCard,
		SelectionStatus: hand.FreeTrumpSelectionStatus,
		StartingPlayer:  hand.StartingPlayer,
		PreviousTrick:   dumpTrick(hand.PreviousTrick),
	}
}

func dumpInProgressHand(hand *game.Hand) *InProgressHandDump {
	trick := hand.GetTrick()
	if trick == nil {
		return nil
	}

	return &InProgressHandDump{
		State:         hand.GetState(),
		Trump:         hand.GetTrump(),
		Trick:         *dumpTrick(trick),
		PreviousTrick: dumpTrick(hand.PreviousTrick),
		Totals:        hand.Totals,
	}
}

func dumpTrick(trick *game.Trick) *TrickDump {
	if trick == nil {
		return nil
	}

	return &TrickDump{
		PlayedCards:    trick.GetTableCards(),
		StartingPlayer: trick.StartingPlayer,
	}
}
