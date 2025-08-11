package room

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand/v2"
	"sync"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server/internal/room/gamecmd"
)

type Room struct {
	Id    string
	Game  game.BeloteGame
	Users map[string]UserData

	started bool

	mu sync.Mutex
}

type messageSender interface {
	SendMessage(msg []byte) error
}

type UserData struct {
	playerId game.PlayerId
	team     game.TeamId
	conn     messageSender
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
		mu:      sync.Mutex{},
	}
}

func (r *Room) Join(userId string, conn messageSender) error {
	r.mu.Lock()
	defer r.mu.Unlock()

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
		playerId: game.NoPlayerId,
		team:     game.Team1,
		conn:     conn,
	}

	return nil
}

func (r *Room) ChooseTeam(userId string, team game.TeamId) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Users[userId]; !ok {
		return ErrPlayerNotFound
	}

	if r.started {
		return ErrGameAlreadyStarted
	}

	r.Users[userId] = UserData{
		playerId: r.Users[userId].playerId,
		team:     team,
		conn:     r.Users[userId].conn,
	}

	return nil
}

func (r *Room) PlayTurn(userId string, gameCmd gamecmd.PlayableCmd) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return gameCmd.PlayTurnAs(r.Users[userId].playerId, &r.Game)
}

func (r *Room) StartGame() error {
	r.mu.Lock()
	defer r.mu.Unlock()

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

func (r *Room) BroadcastState() {
	r.mu.Lock()
	defer r.mu.Unlock()

	for userId, user := range r.Users {
		userState, err := r.DumpUserState(userId)
		if err != nil {
			log.Println("Error dumping user state:", err)
			continue
		}

		var userStateJson []byte
		userStateJson, err = json.Marshal(userState)
		if err != nil {
			log.Println("Error marshalling user state to JSON:", err)
			continue
		}

		go user.conn.SendMessage(userStateJson)
	}
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
