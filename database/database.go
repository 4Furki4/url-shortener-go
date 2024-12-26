package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./url_shortener.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite database: %v", err)
	}

	// Create table if it doesn't exist
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            short_url TEXT NOT NULL UNIQUE,
            original_url TEXT NOT NULL
        );
    `
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func SaveURL(shortURL, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := DB.Exec(query, shortURL, originalURL)
	return err
}

func GetURL(shortURL string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_url = ?`
	var originalURL string
	err := DB.QueryRow(query, shortURL).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", err
	}
	return originalURL, err
}
