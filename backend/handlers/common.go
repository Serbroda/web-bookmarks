package handlers

import (
	"backend/security"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handleError(ctx echo.Context, err error, code int) error {
	return ctx.String(code, err.Error())
}

func getAuthenticatedUser(ctx echo.Context) (security.Authentication, error) {
	auth, ok := ctx.Get(security.ContextKeyAuthentication).(security.Authentication)
	if !ok {
		return auth, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	return auth, nil
}
