package db

import (
	"database/sql"
	"embed"
	"github.com/Serbroda/ragbag/app/pkg/sqlc"
	"sync"
)

var once sync.Once

var (
	Con     *sql.DB
	Queries *sqlc.Queries
)

func OpenAndConfigure(driver, source string, migrations embed.FS, migrationsDir string) {
	db := OpenConnection(driver, source)
	Migrate(db, driver, migrations, migrationsDir)
	InitQueries(db)
}

func OpenConnection(driver, source string) *sql.DB {
	once.Do(func() {
		db, err := connectDatabase(driver, source)
		if err != nil {
			panic("Failed to open database: " + err.Error())
		}
		Con = db
	})
	return Con
}

func InitQueries(db *sql.DB) *sqlc.Queries {
	Queries = sqlc.New(db)
	return Queries
}
