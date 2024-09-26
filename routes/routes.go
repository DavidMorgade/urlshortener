package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/shorten", createShortURLOnDatabase)
	server.GET("/:shortURL", redirectURL)
}
