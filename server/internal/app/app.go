package app

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/room"
	"github.com/los-dogos-studio/gurian-belote/server/internal/token"
	"github.com/los-dogos-studio/gurian-belote/server/internal/userconn"
)

type User struct {
	Id   string
	conn *userconn.UserConn
}

type App struct {
	users        map[string]*User
	roomManager  room.RoomManager
	tokenManager token.TokenManager
	mu           sync.Mutex
}

func NewApp() App {
	return App{
		users:        make(map[string]*User),
		roomManager:  room.NewRoomManager(),
		tokenManager: *token.NewTokenManager(),
		mu:           sync.Mutex{},
	}
}

func (app *App) HandleNewConnection(token string, userName string, ws *websocket.Conn) {
	app.mu.Lock()
	defer app.mu.Unlock()

	conn := userconn.NewUserConn(
		token,
		userName,
		&app.roomManager,
		ws,
	)

	user := &User{Id: token, conn: conn}
	app.users[token] = user

	go conn.Serve()
}

func (app *App) GenerateToken(userName string) string {
	return app.tokenManager.GenerateToken(userName)
}

func (app *App) HandleReconnection(token, roomId string, ws *websocket.Conn) {
	app.mu.Lock()
	defer app.mu.Unlock()

	if !app.validateReconnection(token, roomId) {
		ws.WriteMessage(websocket.TextMessage, []byte(`{"error":"Invalid session"}`))
		ws.Close()
		return
	}

	userSession, _ := app.tokenManager.GetToken(token)

	if oldUser, ok := app.users[token]; ok {
		oldUser.conn.Close()
	}

	conn := userconn.NewUserConn(
		token,
		userSession.UserName,
		&app.roomManager,
		ws,
	)

	room, _ := app.roomManager.GetRoom(roomId)

	conn.Room = room
	room.UpdateUserConnection(token, conn)

	user := &User{Id: token, conn: conn}
	app.users[token] = user

	go conn.Serve()

	userState, err := room.DumpUserState(token)
	if err == nil {
		userStateJson, err := json.Marshal(userState)
		if err == nil {
			conn.SendMessage(userStateJson)
		}
	}
}

func (app *App) validateReconnection(token, roomId string) bool {
	_, tokenExists := app.tokenManager.GetToken(token)
	if !tokenExists {
		return false
	}

	room, roomExists := app.roomManager.GetRoom(roomId)
	if !roomExists {
		return false
	}

	_, userInRoom := room.Users[token]
	return userInRoom
}
