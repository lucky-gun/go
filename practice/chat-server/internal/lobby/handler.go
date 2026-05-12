package lobby

import (
	"errors"

	"chat-server/internal/model"
)

type EventHandler struct {
	RoomService *RoomService
}

func NewEventHandler(roomService *RoomService) *EventHandler {
	return &EventHandler{
		RoomService: roomService,
	}
}

func (h *EventHandler) HandleEvent(
	user *model.User,
	payload *model.ChatPayload,
) error {

	switch payload.Type {

	case "create_room":
		_, err := h.RoomService.CreateRoom(
			payload.RoomName,
			user.ID,
		)

		return err

	case "join_room":
		return h.RoomService.JoinRoom(
			payload.RoomID,
			user,
		)

	case "leave_room":
		return h.RoomService.LeaveRoom(
			payload.RoomID,
			user.ID,
		)

	default:
		return errors.New("unknown event type")
	}
}
