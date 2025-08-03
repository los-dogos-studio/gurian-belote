package userconn

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/server/internal/room/gamecmd"
)

type PlayTurnCmd struct {
	Move json.RawMessage
}

func NewPlayTurnCmd(msg []byte) (Cmd, error) {
	playTurnCmd := PlayTurnCmd{}
	err := json.Unmarshal(msg, &playTurnCmd)
	if err != nil {
		return nil, err
	}

	if playTurnCmd.Move == nil {
		return nil, ErrInvalidCmdParams
	}

	return &playTurnCmd, nil
}

func (c *PlayTurnCmd) HandleCommand(context *CmdContext) error {
	user := context.user

	roomCmd, err := gamecmd.NewGameCmdFromJson(c.Move)
	if err != nil {
		return err
	}

	if user.Room == nil {
		return ErrUserNotInRoom
	}

	return user.Room.PlayTurn(user.UserId, roomCmd)
}
