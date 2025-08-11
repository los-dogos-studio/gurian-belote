package gamecmd

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type SelectTrumpCommand struct {
	Suit *game.Suit
}

const SelectTrumpCmdType = "selectTrump"

func (c *SelectTrumpCommand) PlayTurnAs(playerId game.PlayerId, game *game.BeloteGame) error {
	return game.SelectTrump(playerId, c.Suit)
}

func newSelectTrumpCommand(cmdBytes []byte) (*SelectTrumpCommand, error) {
	selectTrumpCmd := &SelectTrumpCommand{}

	// TODO: Validate
	err := json.Unmarshal(cmdBytes, selectTrumpCmd)
	if err != nil {
		return nil, err
	}

	return selectTrumpCmd, nil
}
