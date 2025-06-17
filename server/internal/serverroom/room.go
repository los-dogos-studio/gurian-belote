package serverroom

import (
	"errors"
	"math/rand/v2"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server/internal/serverroom/gamecmd"
)

type Room struct {
	Id string

	Game game.BeloteGame

	Users   map[string]UserData
	started bool
}

type UserData struct {
	playerId game.PlayerId
	team     game.TeamId
}

var (
	ErrRoomFull           = errors.New("room: room is full")
	ErrPlayerNotFound     = errors.New("room: player not found")
	ErrTeamsNotBalanced   = errors.New("room: teams are not balanced")
	ErrGameAlreadyStarted = errors.New("room: game already started")
)

func NewRoom(id string) *Room {
	return &Room{
		Id:      id,
		Game:    game.NewBeloteGame(),
		Users:   make(map[string]UserData),
		started: false,
	}
}

func (r *Room) Join(userId string) error {
	if _, ok := r.Users[userId]; ok {
		return nil
	}

	if r.started {
		return ErrGameAlreadyStarted
	}

	if len(r.Users) >= game.NUM_PLAYERS {
		return ErrRoomFull
	}

	r.Users[userId] = UserData{
		team: game.Team1,
	}

	return nil
}

func (r *Room) ChooseTeam(userId string, team game.TeamId) error {
	if _, ok := r.Users[userId]; !ok {
		return ErrPlayerNotFound
	}

	if r.started {
		return ErrGameAlreadyStarted
	}

	r.Users[userId] = UserData{
		team: team,
	}

	return nil
}

func (r *Room) PlayTurn(userId string, gameCmd gamecmd.PlayableCmd) error {
	return gameCmd.PlayTurnAs(r.Users[userId].playerId, &r.Game)
}

func (r *Room) StartGame() error {
	if r.started {
		return ErrGameAlreadyStarted
	}

	err := r.assignPlayerIds()
	if err != nil {
		return err
	}

	r.Game.Start()
	return nil
}

func (r *Room) assignPlayerIds() error {
	if err := r.checkBalance(); err != nil {
		return err
	}

	used := make(map[game.PlayerId]bool)

	for user, userData := range r.Users {
		if userData.team == game.Team1 {
			if used[game.Player3] || (!used[game.Player1] && (rand.IntN(2) == 0)) {
				userData.playerId = game.Player1
				used[game.Player1] = true
			} else {
				userData.playerId = game.Player3
				used[game.Player3] = true
			}
		} else {
			if used[game.Player4] || (!used[game.Player2] && rand.IntN(2) == 0) {
				userData.playerId = game.Player2
				used[game.Player2] = true
			} else {
				userData.playerId = game.Player4
				used[game.Player4] = true
			}
		}
		r.Users[user] = userData
	}

	return nil
}

func (r *Room) checkBalance() error {
	team1Count, team2Count := 0, 0
	for _, userData := range r.Users {
		if userData.team == game.Team1 {
			team1Count++
		} else {
			team2Count++
		}
	}

	if team1Count != 2 || team2Count != 2 {
		return ErrTeamsNotBalanced
	}

	return nil
}
