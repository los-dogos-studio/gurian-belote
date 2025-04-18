package clicmd

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewReader(os.Stdin)
	command, err := scanner.ReadString('\n')

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
