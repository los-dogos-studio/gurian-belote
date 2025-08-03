package userconn

import (
	"errors"
)

var (
	ErrInvalidCmdType    = errors.New("invalid command type")
	ErrInvalidTeamId     = errors.New("invalid team id")
	ErrInvalidCmdParams  = errors.New("invalid command parameters")
	ErrUserAlreadyInRoom = errors.New("user already in room")
	ErrUserNotInRoom     = errors.New("user not in room")
	ErrRoomNotFound      = errors.New("room not found")
)
