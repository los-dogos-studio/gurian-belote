package gamecmd

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type PlayCardCommand struct {
	Card game.Card
}

const PlayCardCmdType = "playCard"

func newPlayCardCommand(cmdBytes []byte) (*PlayCardCommand, error) {
	playCardCmd := &PlayCardCommand{}

	err := json.Unmarshal(cmdBytes, playCardCmd)
	if err != nil {
		return nil, err
	}

	return playCardCmd, nil
}

func (c *PlayCardCommand) PlayTurnAs(playerId game.PlayerId, game *game.BeloteGame) error {
	return game.PlayCard(playerId, c.Card)
}
