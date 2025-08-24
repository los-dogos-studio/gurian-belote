package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/server/internal/app"
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
	http.HandleFunc("/auth", s.handleAuth)
	http.HandleFunc("/ws", s.handleWs)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection from:", r.RemoteAddr)

	token, roomId, userName := s.extractRequestParams(r)

	if token == "" {
		http.Error(w, "Authentication required", http.StatusUnauthorized)
		return
	}

	isReconnection := roomId != ""

	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	if isReconnection {
		s.state.HandleReconnection(token, roomId, ws)
	} else {
		if userName == "" {
			ws.WriteMessage(websocket.TextMessage, []byte(`{"error":"userName required for new connection"}`))
			ws.Close()
			return
		}
		s.state.HandleNewConnection(token, userName, ws)
	}
}

func (s *Server) extractRequestParams(r *http.Request) (token, roomId, userName string) {
	if cookie, err := r.Cookie("token"); err == nil {
		token = cookie.Value
	}
	roomId = r.URL.Query().Get("roomId")
	userName = r.URL.Query().Get("userName")
	return
}

func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userName := r.URL.Query().Get("userName")
	if userName == "" {
		http.Error(w, "userName required", http.StatusBadRequest)
		return
	}

	token := s.state.GenerateToken(userName)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   12 * 60 * 60,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + token + `"}`))
}
