package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SpaceController struct {
}

func RegisterSpaceController(e *echo.Echo, c SpaceController, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/spaces", c.GetSpaces, middlewares...)
}

// @Summary get spaces.
// @Description get spaces.
// @Tags spaces
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/spaces [get]
func (c *SpaceController) GetSpaces(ctx echo.Context) error {
	// {object} gen.Space
	/*user, err := c.getUser(ctx)
	if err != nil {
		return err
	}
	spaces, err := services2.Service.FindUserSpaces(ctx.Request().Context(), user.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, mappers.MapSpaces(spaces))*/
	return ctx.String(http.StatusInternalServerError, "not implemented")
}
