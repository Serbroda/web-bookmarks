package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) ListBookmarks(ctx echo.Context) error {
	resp := Bookmark{
		Url: "www.google.de",
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (Server) CreateBookmark(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (Server) ListSpaces(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
func (Server) CreateSpace(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
