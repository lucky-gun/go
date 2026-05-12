package lobby

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type JoinRoomRequest struct {
	RoomID string `json:"room_id"`
}

type Roominfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserCount int    `json:"user_count"`
}
