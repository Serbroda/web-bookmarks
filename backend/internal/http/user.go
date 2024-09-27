package http

import (
	"backend/internal/product"
	"backend/internal/security"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UsersHandler struct {
	UserService *product.UserServiceImpl
}

func RegisterUsersHandlers(e *echo.Group, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/me", h.GetMe, middlewares...)
}

func (h *UsersHandler) GetMe(ctx echo.Context) error {
	auth, err := security.GetAuthentication(ctx)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(http.StatusOK, auth)
}
