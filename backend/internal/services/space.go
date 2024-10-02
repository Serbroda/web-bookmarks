package services

import "backend/internal/sqlc"

type SpaceService struct {
	queries *sqlc.Queries
}

func NewSpaceService(queries *sqlc.Queries) *SpaceService {
	return &SpaceService{queries: queries}
}

func (s *SpaceService) CreateSpace(authenticated sqlc.User, space *sqlc.Space) error {
	if space.OwnerID < 1 {
		space.OwnerID = authenticated.ID
	}
	return nil
}
