package db

import (
	"database/sql"
	"fmt"
	"time"
)

func Connect(databaseString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)

	db.SetMaxIdleConns(25)

	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database unreachable: %w", err)
	}

	return db, err
}
