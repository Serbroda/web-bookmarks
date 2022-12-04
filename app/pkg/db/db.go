package db

import (
	"database/sql"
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

var (
	DB *sqlx.DB
)

func ConnectAndMigrate(driver, source string, migrations embed.FS, migrationsDir string) {
	db := Connect(driver, source)
	Migrate(db.DB, migrations, migrationsDir)
}

func Connect(driver, source string) *sqlx.DB {
	once.Do(func() {
		db, err := sqlx.Connect(driver, source)
		if err != nil {
			panic("Failed to open database: " + err.Error())
		}
		DB = db
	})
	return DB
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
