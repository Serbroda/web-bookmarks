package db

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var once sync.Once

var (
	DB *sqlx.DB
)

func OpenConnection(driver, source string) *sqlx.DB {
	once.Do(func() {
		db, err := connectDatabase(driver, source)
		if err != nil {
			panic("Failed to open database: " + err.Error())
		}
		DB = db
	})
	return DB
}
