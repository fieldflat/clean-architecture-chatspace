package request

type MessagePostRequest struct {
	RoomID  string `json:"room_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
