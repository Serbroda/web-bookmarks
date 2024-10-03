package dto

import (
	"backend/internal/sqlc"
)

type SpaceDto struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func SpaceDtoFromSpace(space sqlc.Space) SpaceDto {
	return SpaceDto{
		ID:          space.ID,
		Name:        space.Name,
		Description: space.Description,
	}
}

func SpaceDtosFromSpaces(spaces []sqlc.Space) []SpaceDto {
	dtos := make([]SpaceDto, len(spaces))
	for i, space := range spaces {
		dtos[i] = SpaceDtoFromSpace(space)
	}
	return dtos
}
