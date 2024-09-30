package http

import (
	"backend/internal/dto"
	"backend/internal/security"
	"backend/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UsersHandler struct {
	UserService *services.UserServiceImpl
}

func RegisterUsersHandlers(e *echo.Group, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/me", h.GetMe, middlewares...)
}

func (h *UsersHandler) GetMe(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}
	user, err := h.UserService.GetById(auth.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, dto.UserDto{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
}
