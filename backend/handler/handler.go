package handler

import "github.com/gin-gonic/gin"

type RoomHandler interface {
	GetHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
}

type MessageHandler interface {
	GetsHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
}
