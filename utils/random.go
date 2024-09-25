package utils

import "math/rand"

func GenerateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortURL := make([]byte, 9)

	for i := range shortURL {

		shortURL[i] = charset[rand.Intn(len(charset))]

	}

	return string(shortURL)

}
