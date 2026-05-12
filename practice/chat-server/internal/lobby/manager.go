package lobby

import (
	"errors"
	"sync"

	"chat-server/internal/model"
)

type LobbyManager struct {
	Rooms map[string]*model.Room
	Mu    sync.RWMutex
}

func NewLobbyManager() *LobbyManager {
	return &LobbyManager{
		Rooms: make(map[string]*model.Room),
	}
}

func (l *LobbyManager) AddRoom(room *model.Room) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	if _, exists := l.Rooms[room.ID]; exists {
		return errors.New("room already exists")
	}

	l.Rooms[room.ID] = room

	return nil
}

func (l *LobbyManager) DeleteRoom(roomID string) {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	delete(l.Rooms, roomID)
}

func (l *LobbyManager) GetRoom(roomID string) (*model.Room, bool) {
	l.Mu.RLock()
	defer l.Mu.RUnlock()

	room, exists := l.Rooms[roomID]

	return room, exists
}

func (l *LobbyManager) ListRooms() []*model.Room {
	l.Mu.RLock()
	defer l.Mu.RUnlock()

	rooms := make([]*model.Room, 0, len(l.Rooms))

	for _, room := range l.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}
