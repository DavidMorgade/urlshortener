package main

import (
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/urlshortener/db"
	"github.com/urlshortener/routes"
)

var urlMap = make(map[string]string)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}

func shortenURL(context *gin.Context) {

	longURL := context.PostForm("url")

	shortURL := generateShortURL()

	urlMap[shortURL] = longURL

	context.JSON(http.StatusOK, gin.H{"shortURL": shortURL})

}

func redirectURL(c *gin.Context) {

	shortURL := c.Param("shortURL")

	longURL, ok := urlMap[shortURL]

	if !ok {

		c.JSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})

		return

	}

	c.Redirect(http.StatusMovedPermanently, longURL)

}

func generateShortURL() string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortURL := make([]byte, 6)

	for i := range shortURL {

		shortURL[i] = charset[rand.Intn(len(charset))]

	}

	return string(shortURL)

}
