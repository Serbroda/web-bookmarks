package db

import (
	"database/sql"
	"embed"
	"sync"

	"github.com/Serbroda/ragbag/gen"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

var (
	Con     *sql.DB
	Queries *gen.Queries
)

func OpenAndConfigure(driver, source string, migrations embed.FS, migrationsDir string) {
	db := OpenConnection(driver, source)
	Migrate(db, migrations, migrationsDir)
	InitQueries(db)
}

func OpenConnection(driver, source string) *sql.DB {
	once.Do(func() {
		db, err := sql.Open(driver, source)
		if err != nil {
			panic("Failed to open database: " + err.Error())
		}
		Con = db
	})
	return Con
}

func InitQueries(db *sql.DB) *gen.Queries {
	Queries = gen.New(db)
	return Queries
}
