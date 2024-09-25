package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", "shorturl.db")

	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10)

	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {

	createURLTable := `
        CREATE TABLE IF NOT EXISTS urls (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          short_url TEXT NOT NULL,
          real_url TEXT NOT NULL
        );
    `

	_, err := DB.Exec(createURLTable)

	if err != nil {
		panic("Could not create the URL table")
	}

	fmt.Println("Succesfully created the URLs table")

}
