package services

import "backend/internal/sqlc"

type SpaceService struct {
	queries *sqlc.Queries
}

func NewSpaceService(queries *sqlc.Queries) *SpaceService {
	return &SpaceService{queries: queries}
}
