package handlers

import (
	"backend/internal/security"
	"backend/internal/service"
	"github.com/labstack/echo/v4"
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
	return ctx.JSON(http.StatusOK, auth)
}
