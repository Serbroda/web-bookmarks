package http

import (
	"backend/internal"
	"backend/internal/common/slice"
	"backend/internal/product"
	"backend/internal/security"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

type CreateSpaceRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

type SpaceHandler struct {
	ContentService *product.ContentService
}

func RegisterSpaceHandlers(e *echo.Group, h SpaceHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/spaces", h.CreateSpace, middlewares...)
	e.GET(baseUrl+"/spaces", h.GetSpaces, middlewares...)
	e.GET(baseUrl+"/spaces/:id", h.GetSpaceById, middlewares...)
	e.DELETE(baseUrl+"/spaces/:id", h.DeleteSpace, middlewares...)
	e.GET(baseUrl+"/spaces/:id/pages/tree", h.GetPagesTreeBySpaceId, middlewares...)
}

func (h *SpaceHandler) CreateSpace(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	var payload CreateSpaceRequest
	if err := BindAndValidate(ctx, &payload); err != nil {
		return err
	}

	space := &internal.Space{
		Name:        payload.Name,
		Description: payload.Description,
		OwnerID:     auth.UserId,
	}

	if err := h.ContentService.CreateSpace(context.TODO(), space); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, space)
}

func (h *SpaceHandler) GetSpaces(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}
	spaces, err := h.ContentService.GetSpacesForUser(auth.UserId)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(http.StatusOK, spaces)
}

func (h *SpaceHandler) GetSpaceById(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	space, err := h.ContentService.GetSpaceById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.String(http.StatusNotFound, err.Error())
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err := h.checkPermission(&space, auth, false); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, space)
}

func (h *SpaceHandler) DeleteSpace(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	space, err := h.ContentService.GetSpaceById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.String(http.StatusNotFound, err.Error())
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err := h.checkPermission(&space, auth, true); err != nil {
		return err
	}

	err = h.ContentService.DeleteSpace(context.TODO(), id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(http.StatusOK, "ok")
}

func (h *SpaceHandler) GetPagesTreeBySpaceId(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	space, err := h.ContentService.GetSpaceById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.String(http.StatusNotFound, err.Error())
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if err := h.checkPermission(&space, auth, false); err != nil {
		return err
	}

	pages, err := h.ContentService.GetPagesBySpaceId(context.TODO(), space.ID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	pagesPointers := slice.ConvertToPointerSlice(pages)

	pagesTree := h.ContentService.BuildPageTree(pagesPointers)
	return ctx.JSON(http.StatusOK, pagesTree)
}

func (h *SpaceHandler) checkPermission(space *internal.Space, auth security.Authentication, requireAdmin bool) error {
	if space.OwnerID != auth.UserId && !slice.ContainsFiltered(space.Shared, func(s internal.UserIdWithRole) bool {
		return s.UserID == auth.UserId && (!requireAdmin || s.Role == "ADMIN")
	}) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}
	return nil
}
