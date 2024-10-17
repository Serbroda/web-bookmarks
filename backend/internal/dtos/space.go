package dtos

import (
	"backend/internal/common/slice"
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

func SpaceDtoFromFindSpaceByIdAndUserIdRow(space sqlc2.FindSpaceByIdAndUserIdRow) SpaceDto {
	return SpaceDto{
		ID:          space.ID,
		Name:        space.Name,
		Description: space.Description,
		Role:        &space.Role,
	}
}

func SpaceDtoFromFindSpacesByUserIdRow(spaces []sqlc2.FindSpacesByUserIdRow) []SpaceDto {
	return slice.MapSlice(spaces, func(item sqlc2.FindSpacesByUserIdRow) SpaceDto {
		return SpaceDto{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Role:        &item.Role,
		}
	})
}
