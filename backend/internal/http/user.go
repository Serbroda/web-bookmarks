package http

import (
	"backend/internal/dtos"
	"backend/internal/security"
	"backend/internal/services"
	"database/sql"
	"errors"
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
		return err
	}

	user, err := h.UserService.GetUserById(auth.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dtos.UserDtoFromUser(user))
}
