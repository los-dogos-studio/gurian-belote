package serverroom

import (
	"errors"

	"github.com/los-dogos-studio/gurian-belote/game"
)

// TODO: should this exist in game package?
type GameState string

type TrickDump struct {
	PlayedCards    map[game.PlayerId]game.Card
	TrumpSuit      game.Suit
	StartingPlayer game.PlayerId
}

type StateDump struct {
	RoomId         string
	Players        map[game.PlayerId]string
	Teams          map[game.TeamId][]string
	Trick          *TrickDump
	StartingPlayer *game.PlayerId
	HandState      game.HandState
	GameState      game.GameState
	Scores         map[game.TeamId]int
}

type UserStateDump struct {
	GameState StateDump
	UserId    string
	PlayerId  game.PlayerId
	UserCards []game.Card
}

var (
	ErrUserNotInRoom = errors.New("user not in room")
)

func DumpState(room *Room) StateDump {
	return StateDump{
		RoomId:         room.Id,
		Players:        dumpPlayersMap(room),
		Teams:          dumpTeams(room),
		Trick:          dumpTrick(room),
		StartingPlayer: dumpStartingPlayer(room),
		HandState:      dumpHandState(room),
		GameState:      dumpGameState(room),
		Scores:         dumpScore(room),
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

func dumpTrick(room *Room) *TrickDump {
	hand := room.Game.GetHand()
	if hand == nil {
		return nil
	}

	trick := room.Game.GetHand().GetTrick()
	if trick == nil {
		return nil
	}

	return &TrickDump{
		PlayedCards:    trick.GetTableCards(),
		TrumpSuit:      hand.GetTrump(),
		StartingPlayer: trick.StartingPlayer,
	}
}

func dumpStartingPlayer(room *Room) *game.PlayerId {
	hand := room.Game.GetHand()
	if hand == nil {
		return nil
	}

	startingPlayer := hand.StartingPlayer
	return &startingPlayer
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
