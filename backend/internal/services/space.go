package services

import (
	sqlc2 "backend/internal/db/sqlc"
	"backend/internal/security"
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
	queries *sqlc2.Queries
}

func NewSpaceService(queries *sqlc2.Queries) *SpaceService {
	return &SpaceService{queries: queries}
}

func (s *SpaceService) CreateSpace(auth security.Authentication, space sqlc2.CreateSpaceParams) (sqlc2.Space, error) {
	if space.Visibility == "" {
		space.Visibility = SpaceVisibilityPrivate
	}

	entity, err := s.queries.CreateSpace(context.TODO(), space)
	if err != nil {
		return sqlc2.Space{}, err
	}

	_, err = s.queries.CreateSpaceUser(context.TODO(), sqlc2.CreateSpaceUserParams{
		SpaceID: entity.ID,
		UserID:  auth.UserId,
		Role:    "OWNER",
	})
	if err != nil {
		return sqlc2.Space{}, err
	}

	return entity, nil
}

func (s *SpaceService) GetSpaceById(auth security.Authentication, spaceId int64) (sqlc2.FindSpaceByIdAndUserIdRow, error) {
	space, err := s.queries.FindSpaceByIdAndUserId(context.TODO(), sqlc2.FindSpaceByIdAndUserIdParams{
		UserID: auth.UserId,
		ID:     spaceId,
	})
	if err != nil {
		return sqlc2.FindSpaceByIdAndUserIdRow{}, err
	}

	count, err := s.queries.CountSpacesUsers(context.TODO(), sqlc2.CountSpacesUsersParams{
		SpaceID: space.ID,
		UserID:  auth.UserId,
	})
	if err != nil {
		return sqlc2.FindSpaceByIdAndUserIdRow{}, err
	} else if count == 0 {
		return sqlc2.FindSpaceByIdAndUserIdRow{}, ErrNoPermission
	}

	return space, nil
}

func (s *SpaceService) GetSpacesByUser(auth security.Authentication) ([]sqlc2.FindSpacesByUserIdRow, error) {
	return s.queries.FindSpacesByUserId(context.TODO(), auth.UserId)
}
