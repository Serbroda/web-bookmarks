package mappers

import (
	"github.com/Serbroda/ragbag/app/cmd/rest-server/handlers/restricted"
	"github.com/Serbroda/ragbag/app/pkg/sqlc"
	"github.com/Serbroda/ragbag/app/pkg/utils"
)

func MapSpace(space sqlc.Space) restricted.SpaceDto {
	return restricted.SpaceDto{
		Description: utils.NullStringToString(space.Description),
		Id:          &space.ShortID,
		Name:        &space.Name,
		OwnerId:     &space.OwnerID,
		Visibility:  (*restricted.SpaceVisibility)(&space.Visibility),
	}
}

func MapSpaces(spaces []sqlc.Space) []restricted.SpaceDto {
	var dtos []restricted.SpaceDto
	for _, s := range spaces {
		dtos = append(dtos, MapSpace(s))
	}
	return dtos
}
