package handlers

import (
	"fmt"

	gen "github.com/Serbroda/ragbag/gen/restricted"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RestrictedServerInterfaceImpl struct {
}

// Delete a space
// (DELETE /groups/{groupId})
func (si *RestrictedServerInterfaceImpl) DeleteGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a group
// (GET /groups/{groupId})
func (si *RestrictedServerInterfaceImpl) GetGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a group
// (PATCH /groups/{groupId})
func (si *RestrictedServerInterfaceImpl) UpdateGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List links of a group
// (GET /groups/{groupId}/links)
func (si *RestrictedServerInterfaceImpl) GetLinks(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a link
// (POST /groups/{groupId}/links)
func (si *RestrictedServerInterfaceImpl) CreateLink(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Delete a link
// (DELETE /links/{linkId})
func (si *RestrictedServerInterfaceImpl) DeleteLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a link
// (GET /links/{linkId})
func (si *RestrictedServerInterfaceImpl) GetLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a link
// (PATCH /links/{linkId})
func (si *RestrictedServerInterfaceImpl) UpdateLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List all spaces
// (GET /spaces)
func (si *RestrictedServerInterfaceImpl) GetSpaces(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Subject
	fmt.Println(name)

	panic("not implemented") // TODO: Implement
}

// Create a space
// (POST /spaces)
func (si *RestrictedServerInterfaceImpl) CreateSpace(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Delete a space
// (DELETE /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) DeleteSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a space
// (GET /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) GetSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a space
// (PATCH /spaces/{spaceId})
func (si *RestrictedServerInterfaceImpl) UpdateSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List groups of a space
// (GET /spaces/{spaceId}/groups)
func (si *RestrictedServerInterfaceImpl) GetGroups(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a group
// (POST /spaces/{spaceId}/groups)
func (si *RestrictedServerInterfaceImpl) CreateGroup(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}
