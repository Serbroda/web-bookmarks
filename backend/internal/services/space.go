package services

import (
	"backend/internal/security"
	"backend/internal/sqlc"
	"context"
)

type SpaceVisibility = string

const (
	SpaceVisibilityPublic  SpaceVisibility = "PUBLIC"
	SpaceVisibilityPrivate SpaceVisibility = "PRIVATE"
)

type SpaceService struct {
	queries *sqlc.Queries
}

func NewSpaceService(queries *sqlc.Queries) *SpaceService {
	return &SpaceService{queries: queries}
}

func (s *SpaceService) CreateSpace(authenticated security.Authentication, space sqlc.CreateSpaceParams) (sqlc.Space, error) {
	if space.Visibility == "" {
		space.Visibility = SpaceVisibilityPrivate
	}

	entity, err := s.queries.CreateSpace(context.TODO(), space)
	if err != nil {
		return sqlc.Space{}, err
	}

	_, err = s.queries.CreateSpaceUser(context.TODO(), sqlc.CreateSpaceUserParams{
		SpaceID: entity.ID,
		UserID:  authenticated.UserId,
		Role:    "OWNER",
	})
	if err != nil {
		return sqlc.Space{}, err
	}

	return entity, nil
}

func (s *SpaceService) GetSpaceById(authenticated sqlc.User, spaceId int64) (sqlc.Space, error) {
	return s.queries.FindSpaceById(context.TODO(), spaceId)
}
