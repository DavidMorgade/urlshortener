package models

import (
	"github.com/urlshortener/db"
)

type URL struct {
	ID       int64  `json:"id"`
	ShortURL string `json:"short_url"`
	RealURL  string `json:"real_url"`
}

func (u *URL) SaveURL() error {
	query := "INSERT INTO urls (short_url, real_url) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.ShortURL, u.RealURL)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err

}

func (u *URL) GetRealURL() (realURL string, err error) {
	query := "SELECT real_url FROM urls WHERE short_url = ?"

	row := db.DB.QueryRow(query, u.ShortURL)

	err = row.Scan(&realURL)

	if err != nil {
		return "", err
	}

	return realURL, nil
}
