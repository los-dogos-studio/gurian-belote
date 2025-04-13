package clicmd

import (
	"fmt"
	"strings"
)

type CliCmd interface {
	ToWsMsg() (string, error)
}

type CliCmdParser interface {
	GetFormat() string
	FromInput(input string) (CliCmd, error)
}

var (
	ErrInvalidCmd = fmt.Errorf("invalid command")
)

var AvailableParsers []CliCmdParser = []CliCmdParser{
	&NewRoomCmdParser{},
	&JoinRoomCmdParser{},
	&StartGameCmdParser{},
	&ChooseTeamCmdParser{},
	&PlayCardCmdParser{},
	&AcceptTrumpCmdParser{},
	&SelectTrumpCmdParser{},
}

func ReadCommand() (CliCmd, error) {
	var command string
	_, err := fmt.Scanln(&command)
	if err != nil {
		return nil, err
	}
	command = strings.TrimSpace(command)

	for _, parser := range AvailableParsers {
		cmd, err := parser.FromInput(command)
		if err == nil {
			return cmd, nil
		}
	}

	return nil, ErrInvalidCmd
}
