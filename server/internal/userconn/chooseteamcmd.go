package userconn

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/game"
)

type ChooseTeamCmd struct {
	TeamId game.TeamId
}

func NewChooseTeamCmd(msg []byte) (Cmd, error) {
	chooseTeamCmd := ChooseTeamCmd{}

	err := json.Unmarshal(msg, &chooseTeamCmd)
	if err != nil {
		return nil, err
	}

	return &chooseTeamCmd, nil
}

func (c *ChooseTeamCmd) HandleCommand(context *CmdContext) error {
	user := context.user

	if c.TeamId != game.Team1 && c.TeamId != game.Team2 {
		return ErrInvalidTeamId
	}

	if user.Room == nil {
		return ErrUserNotInRoom
	}

	return user.Room.ChooseTeam(user.UserId, c.TeamId)
}
