package server

import (
	"github.com/Serbroda/bookmark-manager/internal/api"
	"github.com/Serbroda/bookmark-manager/internal/repository"
	"github.com/labstack/echo/v4"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
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

	e := echo.New()

	server := api.NewServer(repo)
	apiGroup := e.Group("/api/v1")
	swagger, _ := api.GetSwagger()
	apiGroup.Use(oapimiddleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(apiGroup, server)

	return e
}
