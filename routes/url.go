package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urlshortener/models"
	"github.com/urlshortener/utils"
)

func createShortURLOnDatabase(context *gin.Context) {

	var url models.URL

	err := context.BindJSON(&url)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = url.SaveURL()

	url.ShortURL = utils.GenerateShortURL()

	fmt.Println("Real URL: ", url.RealURL)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving URL"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"shortURL": url.ShortURL})

}

func generateShortURL() {
	panic("unimplemented")
}

func redirectURL(c *gin.Context) {

	fmt.Println("Redirecting URL")

	shortURL := c.Param("shorturl")

	var url models.URL

	url.ShortURL = shortURL

	realURL, err := url.GetRealURL()

	fmt.Println("Real URL: ", realURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})
		return
	}
	// Ensure the realURL starts with http:// or https://
	if !strings.HasPrefix(realURL, "http://") && !strings.HasPrefix(realURL, "https://") {
		realURL = "http://" + realURL
	}
	// redirect to the real URL, ommiting my domain
	c.Redirect(http.StatusMovedPermanently, realURL)

	fmt.Println("Redirecting to: ", realURL)

}
