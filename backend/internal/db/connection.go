package db

import (
	migrations2 "backend/internal/db/migrations"
	"database/sql"
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var once sync.Once

var (
	DB *sql.DB
)

func OpenAndConfigure(driver, source string, migrations embed.FS, migrationsDir string) *sql.DB {
	db := OpenConnection(driver, source)
	migrations2.Migrate(db, driver, migrations, migrationsDir)
	return db
}

func OpenConnection(driver, source string) *sql.DB {
	once.Do(func() {
		db, err := connectDatabase(driver, source)
		if err != nil {
			panic("Failed to open database: " + err.Error())
		}
		DB = db
	})
	return DB
}

func connectDatabase(driver, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return db, nil
}
