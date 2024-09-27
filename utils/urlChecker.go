package utils

import (
	"net/http"
	"strings"

	"github.com/urlshortener/models"
)

func CheckIfValidURL(url *models.URL) bool {

	if !strings.HasPrefix(url.RealURL, "http://") && !strings.HasPrefix(url.RealURL, "https://") {
		url.RealURL = "http://" + url.RealURL
	}

	_, err := http.Get(url.RealURL)
	if err != nil {
		return false
	}
	return true
}
