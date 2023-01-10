package db

import (
	"embed"
	"github.com/jmoiron/sqlx"
	"sync"
)

var once sync.Once

var (
	DB *sqlx.DB
)

func OpenAndConfigure(driver, source string, migrations embed.FS, migrationsDir string) *sqlx.DB {
	db := OpenConnection(driver, source)
	Migrate(db.DB, driver, migrations, migrationsDir)
	return db
}

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
