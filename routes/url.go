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

	urlExists := url.CheckIfRealURLExists()

	if urlExists {
		shortURL, err := url.GetShortURL()
		fmt.Println("URL Exists:", url.RealURL)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting real URL"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"shortURL": shortURL, "realURL": url.RealURL})
		return
	}

	url.ShortURL = utils.GenerateShortURL()

	urlExists = url.CheckIfShortURLExists()

	for urlExists {
		url.ShortURL = utils.GenerateShortURL()
		urlExists = url.CheckIfShortURLExists()
	}

	err = url.SaveURL()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving URL"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"shortURL": url.ShortURL, "realURL": url.RealURL})

}

func redirectURL(c *gin.Context) {

	fmt.Println("Redirecting URL")

	shortURL := c.Param("shortURL")

	fmt.Println("Short URL: ", shortURL)

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

// route that gets all the URLs from the databas
func getAllURLs(c *gin.Context) {

	urls, err := models.GetAllURLS()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all URLs"})
		return
	}

	c.JSON(http.StatusOK, urls)

}
