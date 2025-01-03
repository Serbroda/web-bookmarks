package server

import (
	"github.com/Serbroda/bookmark-manager/internal/handlers"
	api2 "github.com/Serbroda/bookmark-manager/internal/handlers/api"
	"github.com/Serbroda/bookmark-manager/internal/repository"
	"github.com/Serbroda/bookmark-manager/internal/security"
	"github.com/Serbroda/bookmark-manager/internal/services"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	userService := services.NewUserService(repo)
	spaceService := services.NewSpaceService(repo)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{
		UserService: userService,
	}, "")

	registerApiV1(e, spaceService)

	return e
}

func registerApiV1(e *echo.Echo, spaceService services.SpaceService) {
	server := api2.NewServer(spaceService)

	apiGroup := e.Group("/api/v1")
	swagger, _ := api2.GetSwagger()
	apiGroup.Use(oapimiddleware.OapiRequestValidator(swagger))
	api2.RegisterHandlers(apiGroup, server)

	apiGroup.Use(echojwt.WithConfig(security.CreateJwtConfig()))
}
