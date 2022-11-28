package services

import (
	"context"

	"github.com/Serbroda/ragbag/gen"
)

func (s *Services) CreateSpace(ctx context.Context, params gen.CreateSpaceParams) (gen.Space, error) {
	id, err := s.Queries.CreateSpace(ctx, params)
	if err != nil {
		return gen.Space{}, err
	}
	space, err := s.Queries.FindSpace(ctx, id)
	if err != nil {
		return gen.Space{}, err
	}
	role, err := s.Queries.FindRoleByName(ctx, "OWNER")
	if err != nil {
		return gen.Space{}, err
	}
	s.Queries.InsertUserSpace(ctx, gen.InsertUserSpaceParams{
		UserID:  params.OwnerID,
		SpaceID: id,
		RoleID:  role.ID,
	})
	return space, nil
}

func (s *Services) FindUserSpaces(ctx context.Context, id int64) ([]gen.Space, error) {
	return s.Queries.FindUserSpaces(ctx, id)
}

func (s *Services) FindUserSpace(ctx context.Context, id int64, spaceId string) (gen.Space, error) {
	return s.Queries.FindUserSpace(ctx, gen.FindUserSpaceParams{
		UserID:  id,
		ShortID: spaceId,
	})
}
