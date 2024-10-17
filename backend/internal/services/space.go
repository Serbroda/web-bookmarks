package services

import (
	"backend/internal/security"
	"backend/internal/sqlc"
	"context"
	"errors"
)

type SpaceVisibility = string

const (
	SpaceVisibilityPublic  SpaceVisibility = "PUBLIC"
	SpaceVisibilityPrivate SpaceVisibility = "PRIVATE"
)

var (
	ErrNoPermission = errors.New("no permission")
)

type SpaceService struct {
	queries *sqlc.Queries
}

func NewSpaceService(queries *sqlc.Queries) *SpaceService {
	return &SpaceService{queries: queries}
}

func (s *SpaceService) CreateSpace(auth security.Authentication, space sqlc.CreateSpaceParams) (sqlc.Space, error) {
	if space.Visibility == "" {
		space.Visibility = SpaceVisibilityPrivate
	}

	entity, err := s.queries.CreateSpace(context.TODO(), space)
	if err != nil {
		return sqlc.Space{}, err
	}

	_, err = s.queries.CreateSpaceUser(context.TODO(), sqlc.CreateSpaceUserParams{
		SpaceID: entity.ID,
		UserID:  auth.UserId,
		Role:    "OWNER",
	})
	if err != nil {
		return sqlc.Space{}, err
	}

	return entity, nil
}

func (s *SpaceService) GetSpaceById(auth security.Authentication, spaceId int64) (sqlc.FindSpaceByIdAndUserIdRow, error) {
	space, err := s.queries.FindSpaceByIdAndUserId(context.TODO(), sqlc.FindSpaceByIdAndUserIdParams{
		UserID: auth.UserId,
		ID:     spaceId,
	})
	if err != nil {
		return sqlc.FindSpaceByIdAndUserIdRow{}, err
	}

	count, err := s.queries.CountSpacesUsers(context.TODO(), sqlc.CountSpacesUsersParams{
		SpaceID: space.ID,
		UserID:  auth.UserId,
	})
	if err != nil {
		return sqlc.FindSpaceByIdAndUserIdRow{}, err
	} else if count == 0 {
		return sqlc.FindSpaceByIdAndUserIdRow{}, ErrNoPermission
	}

	return space, nil
}

func (s *SpaceService) GetSpacesByUser(auth security.Authentication) ([]sqlc.FindSpacesByUserIdRow, error) {
	return s.queries.FindSpacesByUserId(context.TODO(), auth.UserId)
}
