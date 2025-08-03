package server

import (
	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/app"
	"log"
	"net/http"
)

type Server struct {
	state    app.App
	upgrader websocket.Upgrader
}

func checkOrigin(r *http.Request) bool {
	return true // TODO
}

func NewServer() *Server {
	return &Server{
		state: app.NewApp(),
		upgrader: websocket.Upgrader{
			CheckOrigin:     checkOrigin,
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/ws", s.handleWs)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection from:", r.RemoteAddr)

	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	s.state.HandleUserConnection(userId, ws)
}
