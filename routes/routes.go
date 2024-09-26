package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/shorten", createShortURLOnDatabase)
	server.GET("/:shortURL", redirectURL)
	server.GET("/urls", getAllURLs)
}
