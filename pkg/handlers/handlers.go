package handlers

import (
	"github.com/Serbroda/ragbag/gen"
	"github.com/labstack/echo/v4"
)

type ServerInterfaceImpl struct {
}

// Delete a space
// (DELETE /groups/{groupId})
func (si *ServerInterfaceImpl) DeleteGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a group
// (GET /groups/{groupId})
func (si *ServerInterfaceImpl) GetGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a group
// (PATCH /groups/{groupId})
func (si *ServerInterfaceImpl) UpdateGroup(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List links of a group
// (GET /groups/{groupId}/links)
func (si *ServerInterfaceImpl) GetLinks(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a link
// (POST /groups/{groupId}/links)
func (si *ServerInterfaceImpl) CreateLink(ctx echo.Context, groupId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Delete a link
// (DELETE /links/{linkId})
func (si *ServerInterfaceImpl) DeleteLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a link
// (GET /links/{linkId})
func (si *ServerInterfaceImpl) GetLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a link
// (PATCH /links/{linkId})
func (si *ServerInterfaceImpl) UpdateLink(ctx echo.Context, linkId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List all spaces
// (GET /spaces)
func (si *ServerInterfaceImpl) GetSpaces(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Create a space
// (POST /spaces)
func (si *ServerInterfaceImpl) CreateSpace(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Delete a space
// (DELETE /spaces/{spaceId})
func (si *ServerInterfaceImpl) DeleteSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Get a space
// (GET /spaces/{spaceId})
func (si *ServerInterfaceImpl) GetSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Update a space
// (PATCH /spaces/{spaceId})
func (si *ServerInterfaceImpl) UpdateSpace(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// List groups of a space
// (GET /spaces/{spaceId}/groups)
func (si *ServerInterfaceImpl) GetGroups(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}

// Create a group
// (POST /spaces/{spaceId}/groups)
func (si *ServerInterfaceImpl) CreateGroup(ctx echo.Context, spaceId gen.IdString) error {
	panic("not implemented") // TODO: Implement
}
