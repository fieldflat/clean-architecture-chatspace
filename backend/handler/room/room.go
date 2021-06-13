package room

import (
	"chatspace/handler/room/request"
	"chatspace/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type RoomHandlerImpl struct {
	usecase usecase.RoomUsecase
}

func NewRoomHandlerImpl(usecase usecase.RoomUsecase) *RoomHandlerImpl {
	return &RoomHandlerImpl{
		usecase: usecase,
	}
}

func (rhi *RoomHandlerImpl) GetHandler(c *gin.Context) {
	room_id := c.Param("id")
	room, err := rhi.usecase.GetRoom(room_id)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(200, room)
}

func (rhi *RoomHandlerImpl) CreateHandler(c *gin.Context) {
	requestBody := request.RoomPostRequest{}
	c.Bind(&requestBody)
	room, err := rhi.usecase.CreateRoom(requestBody.Description, 60)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(200, room)
}
