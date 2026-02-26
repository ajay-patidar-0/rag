package db

import (
	"database/sql"
	"log"
	"os"
)

func NewDb() (*sql.DB, error) {
	connURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect to DB : %v", err)
	}

	return db, nil
}
