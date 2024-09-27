package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BindAndValidate(ctx echo.Context, payload interface{}) error {
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}
	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}
	return nil
}
