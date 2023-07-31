package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func connectDatabase(driver, source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
