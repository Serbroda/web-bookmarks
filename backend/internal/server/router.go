package server

import (
	"github.com/Serbroda/bookmark-manager/internal/api"
	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()
	e := echo.New()

	apiGroup := e.Group("/api/v1")
	api.RegisterHandlers(apiGroup, server)
	return e
}
