package room

import (
	"chatspace/entity"
	"chatspace/repository/room/inmemory/initdata"
)

type RoomInmemoryRepository struct {
	db map[string]*entity.Room
}

func NewRoomInmemoryRepository() *RoomInmemoryRepository {
	return &RoomInmemoryRepository{
		db: initdata.RoomTestDB,
	}
}

func (rir *RoomInmemoryRepository) Get(rid string) (*entity.Room, error) {
	return rir.db[rid], nil
}

func (rir *RoomInmemoryRepository) Create(description string, minute int) (*entity.Room, error) {
	room := entity.NewRoom(description, minute)
	rir.db[room.ID] = room
	return room, nil
}
