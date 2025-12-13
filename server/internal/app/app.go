package app

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/room"
	"github.com/los-dogos-studio/gurian-belote/server/internal/userconn"
)

type User struct {
	Id        string
	conn      *userconn.UserConn
	sessionId string
}

type App struct {
	users       map[string]*User
	roomManager room.RoomManager
	mu          sync.Mutex
}

func NewApp() App {
	return App{
		users:       make(map[string]*User),
		roomManager: room.NewRoomManager(),
		mu:          sync.Mutex{},
	}
}

func (app *App) HandleUserConnection(
	userId string,
	sessionId string,
	ws *websocket.Conn) {
	app.mu.Lock()
	defer app.mu.Unlock()

	user := app.users[userId]

	if user != nil {
		oldUserConn := user.conn

		userRoom := oldUserConn.Room
		oldUserConn.Close()

		user.conn = userconn.NewUserConn(
			userId,
			userRoom,
			&app.roomManager,
			ws,
			sessionId)
		if userRoom != nil {
			userRoom.UpdateUserConnection(userId, user.conn)
		}
	} else {
		user = &User{
			Id:        userId,
			sessionId: sessionId,
			conn: userconn.NewUserConn(
				userId,
				nil,
				&app.roomManager,
				ws,
				sessionId),
		}

		app.users[userId] = user
	}

	go user.conn.Serve()
}
