package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/server"
)

type StartGameCmd struct{}

type StartGameCmdParser struct{}

func (p *StartGameCmdParser) GetFormat() string {
	return "startGame"
}

func (p *StartGameCmdParser) FromInput(input string) (CliCmd, error) {
	if input != p.GetFormat() {
		return nil, fmt.Errorf("invalid command format: expected '%s', got '%s'", p.GetFormat(), input)
	}
	return &StartGameCmd{}, nil
}

func (s *StartGameCmd) ToWsMsg() (string, error) {
	serverCmd := server.Cmd{
		Command: "startGame",
	}
	msg, err := json.Marshal(serverCmd)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
