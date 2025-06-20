package gamecmd

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type GameCmd struct {
	Command GameCmdType
}

type GameCmdType string

type PlayableCmd interface {
	PlayTurnAs(playerId game.PlayerId, game *game.BeloteGame) error
}

var (
	ErrInvalidCmdType = errors.New("gamecmd: invalid command type")
)

func NewGameCmdFromJson(data json.RawMessage) (PlayableCmd, error) {
	gameCmd := GameCmd{}
	err := json.Unmarshal(data, &gameCmd)
	if err != nil {
		return nil, err
	}

	log.Printf("Game command: %s", gameCmd.Command)
	switch gameCmd.Command {
	case AcceptTrumpCmdType:
		return newAcceptTrumpCommandFromJson(data)
	case SelectTrumpCmdType:
		return newSelectTrumpCommand(data)
	case PlayCardCmdType:
		return newPlayCardCommand(data)
	}
	return nil, ErrInvalidCmdType
}
