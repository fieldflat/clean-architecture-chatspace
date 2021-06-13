package initdata

import (
	"chatspace/entity"
	"time"
)

var RoomTestDB = map[string]*entity.Room{}
var JST, _ = time.LoadLocation("Asia/Tokyo")

func init() {
	RoomTestDB = map[string]*entity.Room{
		"test_room_id": {
			ID:           "test_id",
			Description:  "test_description",
			CreatedAt:    time.Date(2021, 1, 1, 0, 0, 0, 0, JST),
			ExpirationAt: time.Date(2021, 1, 2, 0, 0, 0, 0, JST),
		},
	}
}
