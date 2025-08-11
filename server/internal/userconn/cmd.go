package userconn

import "github.com/los-dogos-studio/gurian-belote/server/internal/room"

type CmdContext struct {
	user        *UserConn
	roomManager *room.RoomManager
}

type Cmd interface {
	HandleCommand(context *CmdContext) error
}
