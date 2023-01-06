package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func connectDatabase(driver, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
