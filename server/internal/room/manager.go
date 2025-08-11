package room

import (
	"strconv"
	"sync"
)

type RoomManager struct {
	rooms map[string]*Room
	idGen *RoomIdGenerator

	mu sync.Mutex
}

func NewRoomManager() RoomManager {
	return RoomManager{
		rooms: make(map[string]*Room),
		idGen: NewRoomIdGenerator(1),
	}
}

func (m *RoomManager) CreateRoom() *Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	roomId := m.idGen.getNextRoomId()
	room := NewRoom(strconv.Itoa(roomId))
	m.rooms[room.Id] = room
	return room
}

func (m *RoomManager) GetRoom(roomId string) (*Room, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	room, exists := m.rooms[roomId]
	return room, exists
}

func (m *RoomManager) DeleteRoom(roomId string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.rooms, roomId)
}

type RoomIdGenerator struct {
	roomIdCounter int
}

func NewRoomIdGenerator(startingId int) *RoomIdGenerator {
	return &RoomIdGenerator{
		roomIdCounter: startingId,
	}
}

func (g *RoomIdGenerator) getNextRoomId() int {
	id := g.roomIdCounter
	g.roomIdCounter++
	return id
}
