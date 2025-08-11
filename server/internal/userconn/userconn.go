package userconn

import (
	"log"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/room"
)

type UserConn struct {
	UserId      string
	Room        *room.Room
	Open        *atomic.Bool
	roomManager *room.RoomManager
	ws          *websocket.Conn
}

func NewUserConn(
	userId string,
	roomManager *room.RoomManager,
	ws *websocket.Conn,
) *UserConn {
	open := &atomic.Bool{}
	open.Store(true)

	return &UserConn{
		UserId:      userId,
		Room:        nil,
		Open:        open,
		roomManager: roomManager,
		ws:          ws,
	}
}

func (c *UserConn) Serve() {
	defer c.ws.Close()
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil || !c.Open.Load() {
			break
		}

		cmd, err := ParseCmd(msg)
		if err != nil {
			// TODO: handle error
			continue
		}

		cmdContext := CmdContext{
			user:        c,
			roomManager: c.roomManager,
		}

		err = cmd.HandleCommand(&cmdContext)
		if err != nil {
			log.Println(err) // FIXME
			continue
		}

		if c.Room != nil {
			c.Room.BroadcastState()
		}
	}
}

func (c *UserConn) SendMessage(msg []byte) error {
	if !c.Open.Load() {
		return nil
	}

	err := c.ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Error sending message:", err)
		return err
	}
	return nil
}

func (c *UserConn) Close() {
	if c.Open.CompareAndSwap(true, false) {
		c.ws.Close()
	}
}
