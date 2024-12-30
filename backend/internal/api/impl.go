package api

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) Server {
	return Server{
		repository: repository,
	}
}

func (Server) ListBookmarks(ctx echo.Context) error {
	resp := Bookmark{
		Url: "www.google.de",
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s Server) CreateBookmark(ctx echo.Context) error {
	bookmark, err := s.repository.CreateBookmark(context.Background(), models.Bookmark{
		URL:   "www.heise.de",
		Title: "Heise",
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

func (Server) ListSpaces(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
func (Server) CreateSpace(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
