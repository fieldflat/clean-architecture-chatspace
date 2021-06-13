package entity

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID           string
	Description  string
	CreatedAt    time.Time
	ExpirationAt time.Time
}

func NewRoom(description string, minute int) *Room {
	rid, err := generateRoomID()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Room{
		ID:           rid,
		Description:  description,
		CreatedAt:    time.Now(),
		ExpirationAt: time.Now().Add(time.Duration(minute) * time.Minute),
	}
}

func generateRoomID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("error: generateRoomID")
	}
	return u.String(), nil
}
