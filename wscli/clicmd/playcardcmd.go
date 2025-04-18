package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server"
)

type PlayCardCmd struct {
	Command string
	Card    game.Card
}

type PlayCardCmdParser struct{}

func (p *PlayCardCmdParser) GetFormat() string {
	return "playCard %s %s"
}

func (p *PlayCardCmdParser) FromInput(input string) (CliCmd, error) {
	// TODO: check if validation is needed
	playCardCmd := PlayCardCmd{
		Command: "playCard",
		Card:    game.Card{},
	}
	_, err := fmt.Sscanf(input, p.GetFormat(), &playCardCmd.Card.Suit, &playCardCmd.Card.Rank)
	if err != nil {
		return nil, err
	}
	return &playCardCmd, nil
}

func (p *PlayCardCmd) ToWsMsg() (string, error) {
	moveMsg, err := json.Marshal(p)
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
