package dtos

import (
	sqlc2 "backend/internal/db/sqlc"
)

type SpaceDto struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Role        *string `json:"role"`
}

func SpaceDtoFromSpace(space sqlc2.Space) SpaceDto {
	return SpaceDto{
		ID:          space.ID,
		Name:        space.Name,
		Description: space.Description,
	}
}

func SpaceDtosFromSpaces(spaces []sqlc2.Space) []SpaceDto {
	dtos := make([]SpaceDto, len(spaces))
	for i, space := range spaces {
		dtos[i] = SpaceDtoFromSpace(space)
	}
	return dtos
}

func SpaceDtoFromFindSpaceByIdAndUserIdRow(space sqlc2.FindSpaceByIdAndUserIdRow) SpaceDto {
	return SpaceDto{
		ID:          space.ID,
		Name:        space.Name,
		Description: space.Description,
		Role:        &space.Role,
	}
}

func SpaceDtoFromFindSpacesByUserIdRow(spaces []sqlc2.FindSpacesByUserIdRow) []SpaceDto {
	dtos := make([]SpaceDto, len(spaces))
	for i, space := range spaces {
		dtos[i] = SpaceDto{
			ID:          space.ID,
			Name:        space.Name,
			Description: space.Description,
			Role:        &space.Role,
		}
	}
	return dtos
}
