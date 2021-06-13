package initdata

import (
	"chatspace/entity"
	"time"
)

var MessageTestDB = map[string][]*entity.Message{}
var JST, _ = time.LoadLocation("Asia/Tokyo")

func init() {
	MessageTestDB = map[string][]*entity.Message{
		"test_room_id": {
			{
				ID:        "init_mid",
				RoomID:    "room_id",
				UserID:    "user_id",
				Content:   "content",
				CreatedAt: time.Now(),
			},
			{
				ID:        "init_mid2",
				RoomID:    "room_id2",
				UserID:    "user_id2",
				Content:   "content2",
				CreatedAt: time.Now().Add(1 * time.Hour),
			},
		},
	}
}
