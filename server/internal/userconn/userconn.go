package userconn

import (
	"encoding/json"
	"errors"
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
	sessionId   string
}

type sessionIdMsg struct {
	SessionId string `json:"sessionId"`
}

var (
	ErrConnectionClosed = errors.New("userconn: connection is closed")
)

func NewUserConn(
	userId string,
	room *room.Room,
	roomManager *room.RoomManager,
	ws *websocket.Conn,
	sessionId string,
) *UserConn {
	open := &atomic.Bool{}
	open.Store(true)

	return &UserConn{
		UserId:      userId,
		Room:        room,
		Open:        open,
		roomManager: roomManager,
		ws:          ws,
		sessionId:   sessionId,
	}
}

func (c *UserConn) Serve() {
	defer c.ws.Close()

	err := c.ws.WriteJSON(sessionIdMsg{
		SessionId: c.sessionId,
	})
	if err != nil {
		log.Println("Error sending sessionId:", err)
	}

	if c.Room != nil {
		c.Room.BroadcastState()
	}

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil || !c.Open.Load() {
			break
		}

		cmd, err := ParseCmd(msg)
		if err != nil {
			log.Println("Error parsing command:", err)
			errMsg := ErrorMessage{
				Error: "invalid command",
			}
			errMsgJson, err := json.Marshal(errMsg)
			if err != nil {
				log.Println("Error marshaling error message:", err)
				continue
			}
			c.SendMessage(errMsgJson)
			continue
		}

		cmdContext := CmdContext{
			user:        c,
			roomManager: c.roomManager,
		}

		err = cmd.HandleCommand(&cmdContext)
		if err != nil {
			errMsg := ErrorMessage{
				Error: err.Error(),
			}
			errMsgJson, err := json.Marshal(errMsg)
			if err != nil {
				log.Println("Error marshaling error message:", err)
				continue
			}
			c.SendMessage(errMsgJson)
			continue
		}

		if c.Room != nil {
			c.Room.BroadcastState()
		}
	}
}

func (c *UserConn) SendMessage(msg []byte) error {
	if !c.Open.Load() {
		// TODO: Check
		return ErrConnectionClosed
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
