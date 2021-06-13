package usecase

import "chatspace/entity"

// type UserUsecase interface {
// 	GetUser(uid string) (*entity.User, error)
// 	CreateUser(name, email, password string) (*entity.User, error)
// 	UpdateUser(name, email, password string) (*entity.User, error)
// }

type RoomUsecase interface {
	GetRoom(rid string) (*entity.Room, error)
	CreateRoom(description string, minute int) (*entity.Room, error)
}

type MessageUsecase interface {
	GetMessages(room_id string) ([]*entity.Message, error)
	CreateMessage(room_id, user_id, content string) (*entity.Message, error)
}
