package server

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/api"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/repository"
	"github.com/labstack/echo/v4"
	"log"
)

func NewServer() *echo.Echo {
	var repo repository.Repository

	repo, err := repository.NewMongoRepository(
		"mongodb://localhost:27017",
		"mongo-golang-test",
	)

	if err != nil {
		log.Fatalf("connection failed")
	}

	repo.CreateBookmark(context.Background(), models.Bookmark{
		URL:         "www.google.de",
		Title:       "Google",
		Description: "Search Engine",
	})

	server := api.NewServer(repo)
	e := echo.New()

	apiGroup := e.Group("/api/v1")
	api.RegisterHandlers(apiGroup, server)
	return e
}
