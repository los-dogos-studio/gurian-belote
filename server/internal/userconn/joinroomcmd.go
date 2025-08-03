package userconn

import (
	"encoding/json"
)

type JoinRoomCmd struct {
	RoomId string
}

func NewJoinRoomCmd(msg []byte) (Cmd, error) {
	joinRoomCmd := JoinRoomCmd{}

	err := json.Unmarshal(msg, &joinRoomCmd)
	if err != nil {
		return nil, err
	}

	return &joinRoomCmd, nil
}

func (c *JoinRoomCmd) HandleCommand(context *CmdContext) error {
	user := context.user
	roomManager := context.roomManager

	if user.Room != nil {
		return ErrUserAlreadyInRoom
	}

	userRoom, ok := roomManager.GetRoom(c.RoomId)
	if !ok {
		return ErrRoomNotFound
	}

	err := userRoom.Join(user.UserId, user)
	if err != nil {
		return err
	}

	user.Room = userRoom
	return nil
}
