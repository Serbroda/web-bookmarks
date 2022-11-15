package database

import (
	"database/sql"
	"embed"
	"log"
	"sync"

	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/glebarez/sqlite"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

var (
	dbConnection *gorm.DB
	db           *sql.DB
	once         sync.Once
)

type ConnectionOptions struct {
	Name          string
	Migrations    embed.FS
	MigrationsDir string
}

func Connect(options ConnectionOptions) *gorm.DB {
	once.Do(func() {
		gormDb, err := gorm.Open(sqlite.Open(options.Name), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", options.Name, err)
			panic(err)
		}

		migrateGorm(gormDb)
		dbConnection = gormDb
	})
	return dbConnection
}

func GetConnection() *gorm.DB {
	return dbConnection
}

func migrateGoose(options ConnectionOptions) {
	db, err := sql.Open("sqlite", options.Name)
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(options.Migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, options.MigrationsDir); err != nil {
		panic(err)
	}
}

func migrateGorm(g *gorm.DB) {
	g.AutoMigrate(&models.User{}, &models.Space{}, &models.Group{}, &models.Link{})
}
