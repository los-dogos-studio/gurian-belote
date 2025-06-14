package serverroom

import (
	"errors"

	"github.com/los-dogos-studio/gurian-belote/game"
)

// TODO: should this exist in game package?
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
}

func (d *FreeTrumpSelectionHandDump) GetState() game.HandState {
	return d.State
}

func (d *FreeTrumpSelectionHandDump) GetStartingPlayer() game.PlayerId {
	return d.StartingPlayer
}

type InProgressHandDump struct {
	State  game.HandState      `json:"state"`
	Trump  game.Suit           `json:"trump"`
	Trick  TrickDump           `json:"trick"`
	Totals map[game.TeamId]int `json:"totals"`
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

func DumpState(room *Room) StateDump {
	return StateDump{
		RoomId:    room.Id,
		Players:   dumpPlayersMap(room),
		Teams:     dumpTeams(room),
		Hand:      dumpHand(room),
		GameState: dumpGameState(room),
		Scores:    dumpScore(room),
	}
}

func DumpUserState(room *Room, userId string) (UserStateDump, error) {
	_, ok := room.Users[userId]
	if !ok {
		return UserStateDump{}, ErrUserNotInRoom
	}

	state := DumpState(room)
	userCards := dumpUserCards(room, userId)

	return UserStateDump{
		GameState: state,
		UserId:    userId,
		PlayerId:  room.Users[userId].playerId,
		UserCards: userCards,
	}, nil
}

func dumpUserCards(room *Room, userId string) []game.Card {
	user, ok := room.Users[userId]
	if !ok {
		return nil
	}

	hand := room.Game.GetHand()
	if hand == nil {
		return nil
	}

	cards := hand.GetPlayerCards(user.playerId)
	result := make([]game.Card, 0, len(cards))
	for card := range cards {
		result = append(result, card)
	}

	return result
}

func dumpPlayersMap(room *Room) map[game.PlayerId]string {
	players := make(map[game.PlayerId]string)
	for id, user := range room.Users {
		players[user.playerId] = id
	}
	return players
}

func dumpTeams(room *Room) map[game.TeamId][]string {
	if room.Users == nil {
		return nil
	}

	teams := make(map[game.TeamId][]string)

	for id, user := range room.Users {
		teams[user.team] = append(teams[user.team], id)
	}
	return teams
}

func dumpHand(room *Room) HandDump {
	hand := room.Game.GetHand()
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

func dumpTableTrumpSelectionHand(hand *game.Hand) *TableTrumpSelectionHandDump {
	return &TableTrumpSelectionHandDump{
		State:           hand.GetState(),
		TableTrumpCard:  hand.TableTrumpCard,
		SelectionStatus: hand.TableTrumpSelectionStatus,
		StartingPlayer:  hand.StartingPlayer,
	}
}

func dumpFreeTrumpSelectionHand(hand *game.Hand) *FreeTrumpSelectionHandDump {
	return &FreeTrumpSelectionHandDump{
		State:           hand.GetState(),
		TableTrumpCard:  hand.TableTrumpCard,
		SelectionStatus: hand.FreeTrumpSelectionStatus,
		StartingPlayer:  hand.StartingPlayer,
	}
}

func dumpInProgressHand(hand *game.Hand) *InProgressHandDump {
	trick := hand.GetTrick()
	if trick == nil {
		return nil
	}

	return &InProgressHandDump{
		State:  hand.GetState(),
		Trump:  hand.GetTrump(),
		Trick:  dumpTrick(trick),
		Totals: hand.Totals,
	}
}

func dumpTrick(trick *game.Trick) TrickDump {
	return TrickDump{
		PlayedCards:    trick.GetTableCards(),
		StartingPlayer: trick.StartingPlayer,
	}
}

func dumpHandState(room *Room) game.HandState {
	if room.Game.GetHand() == nil {
		return game.HandFinished
	}
	return room.Game.GetHand().GetState()
}

func dumpGameState(room *Room) game.GameState {
	return room.Game.GetState()
}

func dumpScore(room *Room) map[game.TeamId]int {
	scores := make(map[game.TeamId]int)
	for team, score := range room.Game.GetScores() {
		scores[team] = score
	}
	return scores
}
