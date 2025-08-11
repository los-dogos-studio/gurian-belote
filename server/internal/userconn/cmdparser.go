package userconn

import (
	"encoding/json"
	"errors"
)

var (
	ErrMessageUnmarshal = errors.New("failed to unmarshal message")
	ErrUnknownCommand   = errors.New("unknown command")
)

type CmdEnvelope struct {
	Command string `json:"command"`
}

type CmdParser func(content []byte) (Cmd, error)

var cmdParsers = map[string]CmdParser{
	"newRoom":    NewCreateRoomCmd, // TODO: remove this alias
	"createRoom": NewCreateRoomCmd,
	"joinRoom":   NewJoinRoomCmd,
	"chooseTeam": NewChooseTeamCmd,
	"startGame":  NewStartGameCmd,
	"playTurn":   NewPlayTurnCmd,
}

func ParseCmd(msg []byte) (Cmd, error) {
	var env CmdEnvelope

	err := json.Unmarshal(msg, &env)
	if err != nil {
		return nil, ErrMessageUnmarshal
	}

	parser, ok := cmdParsers[env.Command]
	if !ok {
		return nil, ErrUnknownCommand
	}

	return parser(msg)
}
