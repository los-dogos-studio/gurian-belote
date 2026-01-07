package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/app"
	"github.com/los-dogos-studio/gurian-belote/server/internal/session"
)

type Server struct {
	app            app.App
	upgrader       websocket.Upgrader
	sessionManager *session.SessionManager
}

const (
	WsCloseCodeInvalidSession = 4001
)

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

	userIdParam := r.URL.Query().Get("userId")
	if userIdParam == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		http.Error(w, "Failed to upgrade to websocket", http.StatusInternalServerError)
		return
	}

	sessionIdParam := r.URL.Query().Get("sessionId")

	sessionId, err := s.authenticateUser(userIdParam, sessionIdParam)
	if err == session.ErrNotAuthenticated {
		wsmsg := websocket.FormatCloseMessage(WsCloseCodeInvalidSession, "User not authenticated")
		ws.WriteControl(websocket.CloseMessage, wsmsg, time.Now().Add(time.Second))
		ws.Close()
		return
	}

	go s.app.HandleUserConnection(userIdParam, sessionId, ws)
}

func (s *Server) authenticateUser(userId string, sessionId string) (string, error) {
	if sessionId == "" {
		return s.sessionManager.CreateSession(userId)
	}

	err := s.sessionManager.AuthUser(userId, sessionId)
	return sessionId, err
}
