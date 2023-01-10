package db

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func connectDatabase(driver, source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
