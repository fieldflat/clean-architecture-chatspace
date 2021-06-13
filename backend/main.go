package main

import (
	"chatspace/handler/room"
	inmemory_room "chatspace/repository/room/inmemory"
	room_service "chatspace/usecase/room"

	"chatspace/handler/message"
	inmemory_message "chatspace/repository/message/inmemory"
	message_service "chatspace/usecase/message"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	roomInmemoryRepository := inmemory_room.NewRoomInmemoryRepository()
	roomUsecase := room_service.NewRoomService(roomInmemoryRepository)
	roomHandler := room.NewRoomHandlerImpl(roomUsecase)

	messageInmemoryRepository := inmemory_message.NewMessageInmemoryRepository()
	messageUsecase := message_service.NewMessageService(messageInmemoryRepository)
	messageHandler := message.NewMessageHandlerImpl(messageUsecase)

	r.GET("/room/:id", roomHandler.GetHandler)
	r.POST("/room", roomHandler.CreateHandler)
	r.GET("/messages/:room_id", messageHandler.GetsHandler)
	r.POST("/message", messageHandler.CreateHandler)
	r.Run(":8081")
}
