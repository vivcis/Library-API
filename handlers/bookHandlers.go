package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vivcis/library-app/db"
)

type Handler struct {
	DB db.Database
}

func PingHandler(c *gin.Context) {
	//healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
