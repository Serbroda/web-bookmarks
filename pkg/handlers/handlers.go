package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/gen/restricted"
	"github.com/Serbroda/ragbag/pkg/handlers/mappers"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RestrictedServerInterfaceImpl struct {
}

// Delete a link
// (DELETE /links/{linkId})
func (si *RestrictedServerInterfaceImpl) DeleteLink(ctx echo.Context, linkId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a link
// (GET /links/{linkId})
func (si *RestrictedServerInterfaceImpl) GetLink(ctx echo.Context, linkId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a link
// (PATCH /links/{linkId})
func (si *RestrictedServerInterfaceImpl) UpdateLink(ctx echo.Context, linkId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Delete a page
// (DELETE /pages/{pageId})
func (si *RestrictedServerInterfaceImpl) DeletePage(ctx echo.Context, pageId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a page
// (GET /pages/{pageId})
func (si *RestrictedServerInterfaceImpl) GetPage(ctx echo.Context, pageId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a page
// (PATCH /pages/{pageId})
func (si *RestrictedServerInterfaceImpl) UpdatePage(ctx echo.Context, pageId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List links of a page
// (GET /pages/{pageId}/links)
func (si *RestrictedServerInterfaceImpl) GetLinks(ctx echo.Context, pageId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a link
// (POST /pages/{pageId}/links)
func (si *RestrictedServerInterfaceImpl) CreateLink(ctx echo.Context, pageId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List all spaces
// (GET /spaces)
func (si *RestrictedServerInterfaceImpl) GetSpaces(ctx echo.Context) error {
	user, err := si.getUser(ctx)
	if err != nil {
		return err
	}
	spaces, err := services.Service.FindUserSpaces(ctx.Request().Context(), user.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, mappers.MapSpaces(spaces))
}

// Create a space
// (POST /spaces)
func (si *RestrictedServerInterfaceImpl) CreateSpace(ctx echo.Context) error {
	var payload restricted.CreateSpaceJSONRequestBody
	err := ctx.Bind(&payload)
	if err != nil || payload.Name == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := si.getUser(ctx)
	if err != nil {
		return err
	}

	params := gen.CreateSpaceParams{
		OwnerID:     user.ID,
		Name:        payload.Name,
		Description: utils.StringToNullString(payload.Description),
	}

	if payload.Visibility != nil {
		params.Visibility = string(*payload.Visibility)
	} else {
		params.Visibility = "PRIVATE"
	}

	space, err := services.Service.CreateSpace(ctx.Request().Context(), params)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, mappers.MapSpace(space))
}

// Delete a space
// (DELETE /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) DeleteSpace(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a space
// (GET /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) GetSpace(ctx echo.Context, spaceId restricted.IdString) error {
	user, err := si.getUser(ctx)
	if err != nil {
		return err
	}
	space, err := services.Service.FindUserSpace(ctx.Request().Context(), user.ID, spaceId)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, mappers.MapSpace(space))
}

// Update a space
// (PATCH /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) UpdateSpace(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List pages of a space
// (GET /spaces/{spaceId}/pages)
func (si *RestrictedServerInterfaceImpl) GetPages(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a page
// (POST /spaces/{spaceId}/pages)
func (si *RestrictedServerInterfaceImpl) CreatePage(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

func (si *RestrictedServerInterfaceImpl) getUser(ctx echo.Context) (gen.User, error) {
	user := ctx.Get("user").(*jwt.Token)
	if user == nil {
		return gen.User{}, echo.ErrUnauthorized
	}

	claims := user.Claims.(*JwtCustomClaims)
	if claims == nil {
		return gen.User{}, echo.ErrUnauthorized
	}

	id, err := utils.ParseInt64(claims.Subject)
	if err != nil {
		return gen.User{}, echo.ErrInternalServerError
	}
	entity, err := services.Service.FindUser(ctx.Request().Context(), id)
	if err != nil {
		if err == services.ErrUserNotFound {
			return gen.User{}, echo.ErrNotFound
		}
		return gen.User{}, err
	}
	return entity, nil
}
