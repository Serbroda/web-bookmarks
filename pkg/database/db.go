package database

import (
	"log"
	"sync"

	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	dbConnection *gorm.DB
	once         sync.Once
)

type ConnectionOptions struct {
	Name string
}

func Connect(options ConnectionOptions) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(options.Name), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", options.Name, err)
			panic(err)
		}
		db.AutoMigrate(&models.User{}, &models.Group{}, &models.Link{}, &models.GroupSubscription{})
		dbConnection = db
	})
	return dbConnection
}

func GetConnection() *gorm.DB {
	return dbConnection
}
