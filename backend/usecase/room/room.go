package room

import (
	"chatspace/entity"
	"chatspace/repository"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(r repository.RoomRepository) *RoomService {
	return &RoomService{
		repo: r,
	}
}

func (rs *RoomService) GetRoom(rid string) (*entity.Room, error) {
	return rs.repo.Get(rid)
}

func (rs *RoomService) CreateRoom(description string, minute int) (*entity.Room, error) {
	return rs.repo.Create(description, minute)
}
