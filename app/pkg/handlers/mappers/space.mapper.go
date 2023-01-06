package mappers

import (
	"github.com/Serbroda/ragbag/app/gen"
	"github.com/Serbroda/ragbag/app/gen/restricted"
	"github.com/Serbroda/ragbag/app/pkg/utils"
)

func MapSpace(space gen.Space) restricted.SpaceDto {
	return restricted.SpaceDto{
		Description: utils.NullStringToString(space.Description),
		Id:          &space.ShortID,
		Name:        &space.Name,
		OwnerId:     &space.OwnerID,
		Visibility:  (*restricted.SpaceVisibility)(&space.Visibility),
	}
}

func MapSpaces(spaces []gen.Space) []restricted.SpaceDto {
	var dtos []restricted.SpaceDto
	for _, s := range spaces {
		dtos = append(dtos, MapSpace(s))
	}
	return dtos
}
