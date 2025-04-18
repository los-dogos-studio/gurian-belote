package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/server"
)

type AcceptTrumpCmd struct {
	Command  string
	Accepted bool
}

type AcceptTrumpCmdParser struct{}

func (p *AcceptTrumpCmdParser) GetFormat() string {
	return "acceptTrump %t"
}

func (p *AcceptTrumpCmdParser) FromInput(input string) (CliCmd, error) {
	cmd := &AcceptTrumpCmd{
		Command:  "acceptTrump",
		Accepted: false,
	}
	_, err := fmt.Sscanf(input, p.GetFormat(), &cmd.Accepted)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *AcceptTrumpCmd) ToWsMsg() (string, error) {
	moveMsg, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}

	serverCmd := server.PlayTurnCmd{
		Command: "playTurn",
		Move:    moveMsg,
	}
	msg, err := json.Marshal(serverCmd)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
