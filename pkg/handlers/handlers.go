package handlers

import (
	"fmt"
	"net/http"

	"github.com/Serbroda/ragbag/gen/restricted"
	"github.com/golang-jwt/jwt/v4"
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
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Subject
	fmt.Println(name)

	return ctx.JSON(http.StatusOK, claims) // TODO: Implement
}

// Create a space
// (POST /spaces)
func (si *RestrictedServerInterfaceImpl) CreateSpace(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Delete a space
// (DELETE /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) DeleteSpace(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a space
// (GET /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) GetSpace(ctx echo.Context, spaceId restricted.IdString) error {
	panic("not implemented") // TODO: Implement
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
