package handlers

import (
	"backend/internal/security"
	"backend/internal/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

type ContentHandler struct {
	ContentService *service.ContentService
}

func RegisterContentHandlers(e *echo.Group, h ContentHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/spaces", h.GetSpaces, middlewares...)
}

func (h *ContentHandler) GetSpaces(ctx echo.Context) error {
	auth, ok := ctx.Get(security.ContextKeyAuthentication).(security.Authentication)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Unauthorized")
	}

	id, err := bson.ObjectIDFromHex(auth.Subject)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "could not parse subject")
	}
	spaces, err := h.ContentService.GetSpacesByUserId(id)
	return ctx.JSON(http.StatusOK, spaces)
}
