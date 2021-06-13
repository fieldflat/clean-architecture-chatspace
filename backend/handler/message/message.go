package message

import (
	"chatspace/handler/message/request"
	"chatspace/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type MessageHandlerImpl struct {
	usecase usecase.MessageUsecase
}

func NewMessageHandlerImpl(usecase usecase.MessageUsecase) *MessageHandlerImpl {
	return &MessageHandlerImpl{
		usecase: usecase,
	}
}

func (rhi *MessageHandlerImpl) GetsHandler(c *gin.Context) {
	room_id := c.Param("room_id")
	messages, err := rhi.usecase.GetMessages(room_id)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(200, messages)
}

func (rhi *MessageHandlerImpl) CreateHandler(c *gin.Context) {
	requestBody := request.MessagePostRequest{}
	c.Bind(&requestBody)
	message, err := rhi.usecase.CreateMessage(requestBody.RoomID, requestBody.UserID, requestBody.Content)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Internal Server Error")
	}
	c.JSON(200, message)
}
