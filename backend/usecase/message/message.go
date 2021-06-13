package message

import (
	"chatspace/entity"
	"chatspace/repository"
)

type MessageService struct {
	repo repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) *MessageService {
	return &MessageService{
		repo: r,
	}
}

func (ms *MessageService) GetMessages(room_id string) ([]*entity.Message, error) {
	return ms.repo.Gets(room_id)
}

func (ms *MessageService) CreateMessage(room_id, user_id, content string) (*entity.Message, error) {
	return ms.repo.Create(room_id, user_id, content)
}
