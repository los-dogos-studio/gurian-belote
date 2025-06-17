package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/serverroom"
)

type User struct {
	RoomId string
	Id     string

	ws *websocket.Conn
}

type Server struct {
	users map[string]*User
	rooms map[string]*serverroom.Room

	roomIdCounter int

	upgrader websocket.Upgrader

	mu sync.Mutex
}

func NewServer() *Server {
	return &Server{
		users:         make(map[string]*User),
		rooms:         make(map[string]*serverroom.Room),
		roomIdCounter: 0,
		upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
			return true
		}, ReadBufferSize: 1024, WriteBufferSize: 1024},
		mu: sync.Mutex{},
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/ws", s.handleWs)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		// TODO
		return
	}
	defer r.Body.Close() // TODO: Remember why
	log.Println("New connection from:", r.RemoteAddr)

	userId := r.URL.Query().Get("userId")
	if userId == "" {
		// TODO
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: auth
	if oldUser, ok := s.users[userId]; ok {
		oldUser.ws.Close()
		delete(s.users, userId)
	}

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO
		return
	}

	user := &User{Id: userId, ws: ws}
	s.users[userId] = user

	go s.serve(user, ws)
}

func (s *Server) serve(user *User, ws *websocket.Conn) {
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			// TODO
		}

		var cmd json.RawMessage

		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			log.Println(err)
			// TODO
			continue
		}

		err = s.handleCommand(user, cmd)
		if err == nil {
			// TODO: Take it to handleCommand
			s.broadcastState(s.rooms[user.RoomId])
		} else {
			log.Println(err)
		}
	}
}

func (s *Server) broadcastState(room *serverroom.Room) {
	for userId := range room.Users {
		userStateDump, err := serverroom.DumpUserState(room, userId)
		if err != nil {
			log.Println(err)
			continue
		}

		data, err := json.Marshal(userStateDump)
		if err != nil {
			log.Println(err)
			continue
		}

		user := s.users[userId] // Lock?
		if user == nil {
			log.Println("User not found:", userId)
			continue
		}
		user.ws.WriteMessage(1, data)
	}
}
