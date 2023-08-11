package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
}

func RegisterUsersHandlers(e *echo.Echo, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/users", h.GetUsers, middlewares...)
}

func (h *UsersHandler) GetUsers(ctx echo.Context) error {
	token, ok := ctx.Get("user").(*jwt.Token)
	fmt.Println(token)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Missing token")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	return ctx.JSON(http.StatusOK, claims)
}
