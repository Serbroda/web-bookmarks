package database

import (
	"database/sql"
	"embed"

	"github.com/Serbroda/ragbag/gen"
	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBCon   *sql.DB
	Queries *gen.Queries
)

func OpenAndConfigure(driver, source string, migrations embed.FS, migrationsDir string) {
	db := OpenConnection(driver, source)
	Migrate(db, migrations, migrationsDir)
	InitialQueries(db)
}

func OpenConnection(driver, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic("Failed to open database: " + err.Error())
	}
	DBCon = db
	return DBCon
}

func Migrate(db *sql.DB, migrations embed.FS, migrationsDir string) {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		panic(err)
	}
}

func InitialQueries(db *sql.DB) {
	Queries = gen.New(db)
}
