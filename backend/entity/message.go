package entity

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	RoomID    string
	UserID    string
	Content   string
	CreatedAt time.Time
}

func NewMessage(room_id, user_id, content string) *Message {
	mid, err := generateMessageID()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Message{
		ID:        mid,
		RoomID:    room_id,
		UserID:    user_id,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func generateMessageID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("error: generateMessageID")
	}
	return u.String(), nil
}
