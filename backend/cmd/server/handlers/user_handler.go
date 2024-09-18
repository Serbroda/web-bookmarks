package handlers

import (
	"backend/internal/security"
	"backend/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
}

type UsersHandler struct {
	UserService *service.UserService
}

func RegisterUsersHandlers(e *echo.Group, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/me", h.GetMe, append(middlewares)...)
}

func (h *UsersHandler) GetMe(ctx echo.Context) error {
	auth, ok := ctx.Get(security.ContextKeyAuthentication).(security.Authentication)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Unauthorized")
	}
	return ctx.JSON(http.StatusOK, auth)
}
