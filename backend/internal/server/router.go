package server

import (
	"github.com/Serbroda/bookmark-manager/internal/api"
	"github.com/Serbroda/bookmark-manager/internal/repository"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

func NewServer(database *mongo.Database) *echo.Echo {
	var repo repository.Repository

	repo, err := repository.NewMongoRepository(database)
	if err != nil {
		log.Fatalf("connection failed")
	}

	server := api.NewServer(&repo)
	e := echo.New()

	apiGroup := e.Group("/api/v1")
	api.RegisterHandlers(apiGroup, server)
	return e
}
