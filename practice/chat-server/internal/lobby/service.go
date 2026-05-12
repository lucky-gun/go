package lobby

import (
	"errors"
	"time"

	"chat-server/internal/model"

	"github.com/google/uuid"
)

type RoomService struct {
	Lobby *LobbyManager
}

func NewRoomService(lobby *LobbyManager) *RoomService {
	return &RoomService{
		Lobby: lobby,
	}
}

func (s *RoomService) CreateRoom(
	name string,
	ownerID string,
) (*model.Room, error) {

	room := &model.Room{
		ID:        uuid.NewString(),
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now(),
	}

	err := s.Lobby.AddRoom(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) RemoveRoom(roomID string) error {
	_, exists := s.Lobby.GetRoom(roomID)

	if !exists {
		return errors.New("room not found")
	}

	s.Lobby.DeleteRoom(roomID)

	return nil
}

func (s *RoomService) JoinRoom(roomID string, user *model.User) error {
	room, exists := s.Lobby.GetRoom(roomID)

	if !exists {
		return errors.New("room not found")
	}

	room.Users[user.ID] = user

	return nil
}

func (s *RoomService) LeaveRoom(roomID string, userID string) error {
	room, exists := s.Lobby.GetRoom(roomID)

	if !exists {
		return errors.New("room not found")
	}

	delete(room.Users, userID)

	if len(room.Users) == 0 {
		s.Lobby.DeleteRoom(roomID)
	}

	return nil
}
