package api

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

type Server struct {
	spaceService services.SpaceService
}

func NewServer(spaceService services.SpaceService) Server {
	return Server{
		spaceService: spaceService,
	}
}

func (s Server) ListSpaces(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
func (s Server) CreateSpace(ctx echo.Context) error {
	var payload CreateSpaceDto
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	entity, err := s.spaceService.CreateSpace(context.TODO(), "", models.Space{
		Name:        payload.Name,
		Description: payload.Description,
		Visibility:  models.SpaceVisibilityPrivate,
		OwnerID:     bson.ObjectID{},
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, entity)
}

func (s Server) DeleteSpace(ctx echo.Context, spaceId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) UpdateSpace(ctx echo.Context, spaceId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) ListBookmarks(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) CreateBookmark(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) DeleteBookmark(ctx echo.Context, bookmarkId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) UpdateBookmark(ctx echo.Context, bookmarkId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) ListCollections(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) CreateCollection(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) DeleteCollection(ctx echo.Context, collectionId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}

func (s Server) UpdateCollection(ctx echo.Context, collectionId Id) error {
	return ctx.String(http.StatusNotImplemented, "Not Implemented")
}
