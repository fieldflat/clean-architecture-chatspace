package repository

import "chatspace/entity"

type RoomRepository interface {
	Get(rid string) (*entity.Room, error)
	Create(description string, minute int) (*entity.Room, error)
}

type MessageRepository interface {
	Gets(room_id string) ([]*entity.Message, error)
	Create(room_id, user_id, content string) (*entity.Message, error)
}
