package handlers

import (
	"backend/models"
	"backend/services"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

type CreateSpaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (j *CreateSpaceRequest) Validate() *ConstraintViolationError {
	var violations ConstraintViolationError

	if len(j.Name) == 0 {
		violations.AddViolation("name", "name must be set")
	}

	if len(violations.Violations) > 0 {
		return &violations
	}
	return nil
}

type SpaceHandler struct {
	ContentService *services.ContentService
}

func RegisterSpaceHandlers(e *echo.Group, h SpaceHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/spaces", h.CreateSpace, middlewares...)
	e.GET(baseUrl+"/spaces", h.GetSpaces, middlewares...)
	e.GET(baseUrl+"/spaces/:id", h.GetSpaceById, middlewares...)
	e.DELETE(baseUrl+"/spaces/:id", h.DeleteSpace, middlewares...)
}

func (h *SpaceHandler) CreateSpace(ctx echo.Context) error {
	auth, err := getAuthenticatedUser(ctx)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}

	var payload CreateSpaceRequest
	if err := ctx.Bind(&payload); err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	if err := payload.Validate(); err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	space := &models.Space{
		Name:        payload.Name,
		Description: payload.Description,
		OwnerID:     auth.UserId,
	}

	if err := h.ContentService.CreateSpace(context.TODO(), space); err != nil {
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusCreated, space)
}

func (h *SpaceHandler) GetSpaces(ctx echo.Context) error {
	auth, err := getAuthenticatedUser(ctx)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}
	spaces, err := h.ContentService.GetSpacesByUserId(auth.UserId)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}
	return ctx.JSON(http.StatusOK, spaces)
}

func (h *SpaceHandler) GetSpaceById(ctx echo.Context) error {
	auth, err := getAuthenticatedUser(ctx)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}

	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	space, err := h.ContentService.GetSpaceById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return handleError(ctx, err, http.StatusNotFound)
		}
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	if space.OwnerID != auth.UserId && !contains(space.Shared, func(s models.UserIdWithRole) bool {
		return s.UserID == auth.UserId
	}) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	return ctx.JSON(http.StatusOK, space)
}

func (h *SpaceHandler) DeleteSpace(ctx echo.Context) error {
	auth, err := getAuthenticatedUser(ctx)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}

	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	space, err := h.ContentService.GetSpaceById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return handleError(ctx, err, http.StatusNotFound)
		}
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	if space.OwnerID != auth.UserId && !contains(space.Shared, func(s models.UserIdWithRole) bool {
		return s.UserID == auth.UserId && s.Role == "ADMIN"
	}) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	err = h.ContentService.DeleteSpace(context.TODO(), id)
	if err != nil {
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	return ctx.String(http.StatusOK, "ok")
}

func contains[T any](slice []T, compare func(T) bool) bool {
	for _, v := range slice {
		if compare(v) {
			return true
		}
	}
	return false
}
