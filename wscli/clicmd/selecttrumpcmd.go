package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server"
)

type SelectTrumpCmd struct {
	Command string
	Trump   *game.Suit
}

type SelectTrumpCmdParser struct{}

func (p *SelectTrumpCmdParser) GetFormat() string {
	return "selectTrump %s"
}

func (p *SelectTrumpCmdParser) FromInput(input string) (CliCmd, error) {
	trumpInput := ""
	_, err := fmt.Sscanf(input, p.GetFormat(), &trumpInput)
	if err != nil {
		return nil, err
	}

	cmd := &SelectTrumpCmd{
		Command: "selectTrump",
		Trump:   nil,
	}

	switch trumpInput {
	case string(game.Spades), string(game.Hearts), string(game.Diamonds), string(game.Clubs):
		suit := game.Suit(trumpInput)
		cmd.Trump = &suit
		return cmd, nil
	case "none":
		cmd.Trump = nil
		return cmd, nil
	default:
		return nil, fmt.Errorf("invalid trump suit")
	}
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
