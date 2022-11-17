package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sync"

	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnection *gorm.DB
	db           *sql.DB
	once         sync.Once
)

type ConnectionOptions struct {
	DbAddress     string
	DbName        string
	DbUser        string
	DbPassword    string
	Migrations    embed.FS
	MigrationsDir string
}

func getDsn(options ConnectionOptions) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", options.DbUser, options.DbPassword, options.DbAddress, options.DbName)
}

func Connect(options ConnectionOptions) *gorm.DB {
	once.Do(func() {
		migrateGoose(options)

		dsn := getDsn(options)
		gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", dsn, err)
			panic(err)
		}

		//migrateGorm(gormDb)
		dbConnection = gormDb
	})
	return dbConnection
}

func GetConnection() *gorm.DB {
	return dbConnection
}

func migrateGoose(options ConnectionOptions) {
	db, err := sql.Open("mysql", getDsn(options))
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(options.Migrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, options.MigrationsDir); err != nil {
		panic(err)
	}
}

func migrateGorm(g *gorm.DB) {
	g.AutoMigrate(&models.User{}, &models.Space{}, &models.Group{}, &models.Link{})
}
