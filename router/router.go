package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vivcis/library-app/handlers"
	"net/http"
	"os"
	"time"
)

var router *gin.Engine

type Router struct {
	ContentType string
	handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

func SetupRouter(h *handlers.Handler) (*gin.Engine, string) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apirouter := router.Group("/api/v1")

	apirouter.GET("/ping", handlers.PingHandler)

	port := os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
