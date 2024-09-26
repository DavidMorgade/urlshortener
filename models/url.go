package models

import (
	"github.com/urlshortener/db"
)

type URL struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"short_url"`
	RealURL  string `json:"real_url"`
}

func GetAllURLS() (urls []URL, err error) {

	// slice of urls
	urls = []URL{}

	query := "SELECT * FROM urls"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var url URL

		err = rows.Scan(&url.ID, &url.ShortURL, &url.RealURL)

		if err != nil {
			return
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func (u *URL) SaveURL() error {
	query := "INSERT INTO urls (short_url, real_url) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(&u.ShortURL, &u.RealURL)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err

}

// method that gets the real URL from the database, using the short URL that is on the url struct pointer
func (u *URL) GetRealURL() (realURL string, err error) {
	query := "SELECT real_url FROM urls WHERE short_url = ?"

	row := db.DB.QueryRow(query, &u.ShortURL)

	err = row.Scan(&realURL)

	if err != nil {
		return "", err
	}

	return realURL, nil
}

func (u *URL) GetShortURL() (shortURL string, err error) {
	query := "SELECT short_url FROM urls WHERE real_url = ?"

	row := db.DB.QueryRow(query, u.RealURL)

	err = row.Scan(&shortURL)

	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (u *URL) CheckIfRealURLExists() bool {
	query := "SELECT id FROM urls WHERE real_url = ?"

	row := db.DB.QueryRow(query, u.RealURL)

	err := row.Scan(&u.ID)

	if err != nil {
		return false
	}

	return true
}
