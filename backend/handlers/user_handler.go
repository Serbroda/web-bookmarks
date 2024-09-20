package handlers

import (
	"backend/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
}

type UsersHandler struct {
	UserService *services.UserService
}

func RegisterUsersHandlers(e *echo.Group, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/me", h.GetMe, middlewares...)
}

func (h *UsersHandler) GetMe(ctx echo.Context) error {
	auth, err := getAuthenticatedUser(ctx)
	if err != nil {
		return handleError(ctx, err, http.StatusUnauthorized)
	}
	return ctx.JSON(http.StatusOK, auth)
}
