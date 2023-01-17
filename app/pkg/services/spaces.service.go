package services

import (
	"context"
	"github.com/Serbroda/ragbag/app/pkg/sqlc"

	"github.com/teris-io/shortid"
)

func (s *Services) CreateSpace(ctx context.Context, params sqlc.CreateSpaceParams) (sqlc.Space, error) {
	if params.ShortID == "" {
		params.ShortID = shortid.MustGenerate()
	}

	id, err := s.Queries.CreateSpace(ctx, params)
	if err != nil {
		return sqlc.Space{}, err
	}
	space, err := s.Queries.FindSpace(ctx, id)
	if err != nil {
		return sqlc.Space{}, err
	}
	role, err := s.Queries.FindRoleByName(ctx, "OWNER")
	if err != nil {
		return sqlc.Space{}, err
	}
	s.Queries.InsertUserSpace(ctx, sqlc.InsertUserSpaceParams{
		UserID:  params.OwnerID,
		SpaceID: id,
		RoleID:  role.ID,
	})
	return space, nil
}

func (s *Services) FindUserSpaces(ctx context.Context, id int64) ([]sqlc.Space, error) {
	return s.Queries.FindUserSpaces(ctx, id)
}

func (s *Services) FindUserSpace(ctx context.Context, id int64, spaceId string) (sqlc.Space, error) {
	return s.Queries.FindUserSpace(ctx, sqlc.FindUserSpaceParams{
		UserID:  id,
		ShortID: spaceId,
	})
}
