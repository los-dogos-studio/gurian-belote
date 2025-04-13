package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server"
)

type SelectTrumpCmd struct {
	Trump game.Suit
}

type SelectTrumpCmdParser struct{}

func (p *SelectTrumpCmdParser) GetFormat() string {
	return "selectTrump %s"
}

func (p *SelectTrumpCmdParser) FromInput(input string) (CliCmd, error) {
	cmd := &SelectTrumpCmd{}
	_, err := fmt.Sscanf(input, p.GetFormat(), &cmd.Trump)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *SelectTrumpCmd) ToWsMsg() (string, error) {
	moveMsg, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}
	msg := server.PlayTurnCmd{
		Command: "playTurn",
		Move:    moveMsg,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(msgBytes), nil
}
