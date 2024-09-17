package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Register(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Not Implemented")
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Not Implemented")
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Not Implemented")
}
