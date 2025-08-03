package app

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/room"
	"github.com/los-dogos-studio/gurian-belote/server/internal/userconn"
)

type User struct {
	Id   string
	conn *userconn.UserConn
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

func (app *App) HandleUserConnection(userId string, ws *websocket.Conn) {
	app.mu.Lock()
	defer app.mu.Unlock()

	// TODO: auth
	if oldUser, ok := app.users[userId]; ok {
		oldUser.conn.Close()
		delete(app.users, userId)
	}

	conn := userconn.NewUserConn(
		userId,
		&app.roomManager,
		ws,
	)

	user := &User{Id: userId, conn: conn}
	app.users[userId] = user

	go conn.Serve()
}
