package gamecmd

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type AcceptTableTrumpCommand struct {
	Accepted bool
}

const AcceptTrumpCmdType = "acceptTrump"

func (c *AcceptTableTrumpCommand) PlayTurnAs(playerId game.PlayerId, game *game.BeloteGame) error {
	return game.AcceptTableTrump(playerId, c.Accepted)
}

func newAcceptTrumpCommandFromJson(cmdBytes []byte) (PlayableCmd, error) {
	acceptTrumpCmd := &AcceptTableTrumpCommand{}

	err := json.Unmarshal(cmdBytes, acceptTrumpCmd)
	if err != nil {
		return nil, err
	}

	return acceptTrumpCmd, nil
}
