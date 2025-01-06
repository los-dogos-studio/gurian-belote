package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/los-dogos-studio/gurian-belote/game"
	"log"
	"net/http"
	"sync"
)

const MaxPlayersPerRoom = 4

type Player struct {
	Id         string
	Conn       *websocket.Conn
	BeloteGame *game.BeloteGame
	RoomId     string
	Username   string
	Team       string // 'red' or 'green'
}

type Room struct {
	Id      string
	Players map[string]*Player
	mu      sync.RWMutex
}

type GameServer struct {
	Rooms     map[string]*Room
	Players   map[string]*Player
	Usernames map[string]bool
	mu        sync.RWMutex
	upgrader  websocket.Upgrader
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

func NewServer() *GameServer {
	return &GameServer{
		Rooms:     make(map[string]*Room),
		Players:   make(map[string]*Player),
		Usernames: make(map[string]bool),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // TODO: restrict origin depending on what akaki will buy
			},
		},
	}
}
func (s *GameServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	playerId := uuid.New().String()

	player := &Player{
		Id:   playerId,
		Conn: conn,
	}

	s.mu.Lock()
	s.Players[playerId] = player
	s.mu.Unlock()

	writeErr := player.Conn.WriteJSON(Message{
		Type: "player_connected",
		Content: map[string]string{
			"player_id": playerId,
		},
	})

	if writeErr != nil {
		log.Printf("Error sending player connected message: %v", writeErr)
		return
	}

	defer s.removePlayer(player)
	go s.listenPlayerMessages(player)
}

func (s *GameServer) joinRoom(player *Player, roomId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, exists := s.Rooms[roomId]
	if !exists {
		room = &Room{
			Id:      roomId,
			Players: make(map[string]*Player),
		}
		s.Rooms[roomId] = room
		log.Printf("Created room %s", roomId)
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	if len(room.Players) >= MaxPlayersPerRoom {

		if len(room.Players) == MaxPlayersPerRoom { // in this case game should be started
			s.broadcastToRoom(roomId, Message{
				Type: "game_action",
				Content: map[string]interface{}{
					"action": "start_game",
				},
			})
			return
		}

		sendErr := player.Conn.WriteJSON(Message{
			Type:    "error",
			Content: "Room is full",
		})
		if sendErr != nil {
			log.Printf("Error sending error message: %v", sendErr)
		}
		closeErr := player.Conn.Close()
		if closeErr != nil {
			log.Printf("Error closing connection: %v", closeErr)
		}
		return
	}

	room.Players[player.Id] = player

	s.broadcastToRoom(roomId, Message{
		Type: "player_joined",
		Content: map[string]string{
			"username":  player.Username,
			"player_id": player.Id,
		},
	})
}

func (s *GameServer) removePlayer(player *Player) {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, exists := s.Rooms[player.RoomId]
	if !exists {
		log.Printf("Room %s does not exist", player.RoomId)
		return
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	delete(room.Players, player.Id)
	closeErr := player.Conn.Close()
	if closeErr != nil {
		log.Printf("Error closing connection: %v", closeErr)
		return
	}

	if len(room.Players) == 0 {
		delete(s.Rooms, player.RoomId)
	} else {
		s.broadcastToRoom(player.RoomId, Message{
			Type: "player_left",
			Content: map[string]string{
				"username":  player.Username,
				"player_id": player.Id,
			},
		})
	}
}

func (s *GameServer) getRoomsList() []map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rooms := make([]map[string]interface{}, 0)
	for roomId, room := range s.Rooms {
		room.mu.RLock()
		rooms = append(rooms, map[string]interface{}{
			"room_id":     roomId,
			"playerCount": len(room.Players),
			"maxPlayers":  MaxPlayersPerRoom,
		})
		room.mu.RUnlock()
	}
	return rooms
}

func (s *GameServer) registerUsername(playerId, username string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Usernames[username] {
		return false
	}

	s.Usernames[username] = true
	if player, exists := s.Players[playerId]; exists {
		player.Username = username
	}
	return true
}

func (s *GameServer) broadcastToRoom(roomId string, msg Message) {
	s.mu.RLock()
	room, exists := s.Rooms[roomId]
	s.mu.RUnlock()

	if !exists {
		log.Printf("Room %s does not exist", roomId)
		return
	}

	room.mu.RLock()
	defer room.mu.RUnlock()

	var beloteGame game.BeloteGame
	var beloteGameStarted bool = false

	if msg.Type == "game_action" {
		content, ok := msg.Content.(map[string]interface{})
		if !ok {
			log.Printf("Invalid content format for game_action")
		}
		action, ok := content["action"].(string)
		if !ok {
			log.Printf("Action not found in content")
		}

		if action == "start_game" {
			beloteGame = game.NewBeloteGame(1001) // TODO: hardcoded value
			beloteGameStarted = true
		}

	}

	for _, player := range room.Players {
		if beloteGameStarted {
			player.BeloteGame = &beloteGame
		}

		err := player.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error broadcasting to player %s: %v", player.Id, err)
		}
	}
}

func (s *GameServer) listenPlayerMessages(player *Player) {
	for {
		var msg Message
		readErr := player.Conn.ReadJSON(&msg)
		if readErr != nil {
			log.Printf("Error reading message from player %s: %v", player.Id, readErr)
			s.removePlayer(player)
			return
		}

		switch msg.Type {
		case "register_username":
			content, ok := msg.Content.(map[string]interface{})
			if !ok {
				log.Printf("Invalid content format for register_username")
				continue
			}
			username, ok := content["username"].(string)
			if !ok {
				log.Printf("Username not found in content")
				continue
			}

			success := s.registerUsername(player.Id, username)
			writeErr := player.Conn.WriteJSON(Message{
				Type: "registration_response",
				Content: map[string]bool{
					"success": success,
				},
			})

			if writeErr != nil {
				log.Printf("Error sending registration response: %v", writeErr)
			}

		case "get_rooms":
			rooms := s.getRoomsList()
			writeErr := player.Conn.WriteJSON(Message{
				Type:    "rooms_list",
				Content: map[string][]map[string]interface{}{"rooms": rooms},
			})

			if writeErr != nil {
				log.Printf("Error sending rooms list: %v", writeErr)
			}

		case "join_room":
			content, ok := msg.Content.(map[string]interface{})
			if !ok {
				log.Printf("Invalid content format for join_room")
				continue
			}
			roomId, ok := content["room_id"].(string)
			if !ok {
				s.joinRoom(player, uuid.New().String())
				continue
			}
			s.joinRoom(player, roomId)

		case "select_team":
			content, ok := msg.Content.(map[string]interface{})
			if !ok {
				log.Printf("Invalid content format for select_team")
				continue
			}
			team, ok := content["team"].(string)
			if !ok {
				log.Printf("Team not found in content")
				continue
			}

			if player.RoomId != "" {
				player.Team = team
				s.broadcastToRoom(player.RoomId, Message{
					Type: "team_selected",
					Content: map[string]string{
						"username":  player.Username,
						"team":      team,
						"player_id": player.Id,
					},
				})
			}

		case "game_action":
			if player.RoomId != "" {
				content, ok := msg.Content.(map[string]interface{})
				if !ok {
					log.Printf("Invalid content format for game_action")
					continue
				}
				action, ok := content["action"].(string)
				if !ok {
					log.Printf("Action not found in content")
					continue
				}

				switch action {
				case "start_game":

				}

			}

			s.broadcastToRoom(player.RoomId, msg)
		}

		//case "chat":
		//	if player.RoomId != "" {
		//		content, ok := msg.Content.(map[string]interface{})
		//		if !ok {
		//			log.Printf("Invalid content format for chat")
		//			continue
		//		}
		//		s.broadcastToRoom(player.RoomId, Message{
		//			Type: "chat",
		//			Content: map[string]interface{}{
		//				"username": player.Username,
		//				"message":  content["message"],
		//				"time":     content["time"],
		//			},
		//		})
		//	}
		//}
	}
}

func (s *GameServer) Start() {
	http.HandleFunc("/ws", s.handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
