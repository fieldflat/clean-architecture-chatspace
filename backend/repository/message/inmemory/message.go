package message

import (
	"chatspace/entity"
	"chatspace/repository/message/inmemory/initdata"
)

type MessageInmemoryRepository struct {
	db map[string][]*entity.Message
}

func NewMessageInmemoryRepository() *MessageInmemoryRepository {
	return &MessageInmemoryRepository{
		db: initdata.MessageTestDB, // key: room_id
	}
}

func (mir *MessageInmemoryRepository) Gets(room_id string) ([]*entity.Message, error) {
	return mir.db[room_id], nil
}

func (mir *MessageInmemoryRepository) Create(room_id, user_id, content string) (*entity.Message, error) {
	message := entity.NewMessage(room_id, user_id, content)
	mir.db[room_id] = append(mir.db[room_id], message)
	return message, nil
}
