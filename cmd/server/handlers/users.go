package handlers

import (
	"net/http"

	"github.com/Serbroda/ragbag/pkg/security"
	"github.com/Serbroda/ragbag/pkg/user"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UserService user.UserService
}

func RegisterUsersHandlers(e *echo.Echo, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/users", h.GetUsers, append(middlewares, security.HasAnyRoleMiddleware("ADMIN"))...)
}

func (h *UsersHandler) GetUsers(ctx echo.Context) error {
	auth, ok := ctx.Get(security.ContextKeyAuthentication).(security.Authentication)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Unauthorized")
	}
	return ctx.JSON(http.StatusOK, auth)
}
