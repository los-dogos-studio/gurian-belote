package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/app"
	"github.com/los-dogos-studio/gurian-belote/server/internal/session"
)

type Server struct {
	app            app.App
	upgrader       websocket.Upgrader
	sessionManager *session.SessionManager
}

func checkOrigin(r *http.Request) bool {
	return true // TODO
}

func NewServer() *Server {
	return &Server{
		app: app.NewApp(),
		upgrader: websocket.Upgrader{
			CheckOrigin:     checkOrigin,
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		sessionManager: session.NewSessionManager(),
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

	sessionId := r.URL.Query().Get("sessionId")
	if sessionId == "" {
		newSessionId, err := s.sessionManager.CreateSession(userId)
		if err == session.ErrNotAuthenticated {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
		sessionId = newSessionId
	} else {
		err := s.sessionManager.AuthUser(userId, sessionId)
		if err == session.ErrNotAuthenticated {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
	}

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		http.Error(w, "Failed to upgrade to websocket", http.StatusInternalServerError)
		return
	}

	go s.app.HandleUserConnection(userId, sessionId, ws)
}
