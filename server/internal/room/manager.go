package room

import (
	"strconv"
	"sync/atomic"
)

type RoomManager struct {
	rooms map[string]*Room
	idGen *RoomIdGenerator
}

func NewRoomManager() RoomManager {
	return RoomManager{
		rooms: make(map[string]*Room),
		idGen: NewRoomIdGenerator(),
	}
}

func (m *RoomManager) CreateRoom() *Room {
	roomId := m.idGen.getNextRoomId()
	room := NewRoom(strconv.Itoa(roomId))
	m.rooms[room.Id] = room
	return room
}

func (m *RoomManager) GetRoom(roomId string) (*Room, bool) {
	room, exists := m.rooms[roomId]
	return room, exists
}

func (m *RoomManager) DeleteRoom(roomId string) {
	delete(m.rooms, roomId)
}

type RoomIdGenerator struct {
	roomIdCounter atomic.Int32
}

func NewRoomIdGenerator() *RoomIdGenerator {
	return &RoomIdGenerator{
		roomIdCounter: atomic.Int32{},
	}
}

func (g *RoomIdGenerator) getNextRoomId() int {
	return int(g.roomIdCounter.Add(1))
}
