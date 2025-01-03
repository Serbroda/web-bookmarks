package services

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/repository"
)

type SpaceService interface {
	CreateSpace(ctx context.Context, userID string, sp models.Space) (models.Space, error)
	GetSpace(ctx context.Context, userID, spaceID string) (models.Space, error)
	GetUserSpaces(ctx context.Context, userID string) ([]models.Space, error)
}

type spaceService struct {
	repo repository.Repository
}

func NewSpaceService(repo repository.Repository) SpaceService {
	return &spaceService{repo: repo}
}

func (s spaceService) CreateSpace(ctx context.Context, userID string, sp models.Space) (models.Space, error) {
	return s.repo.CreateSpace(ctx, sp)
}

func (s spaceService) GetSpace(ctx context.Context, userID, spaceID string) (models.Space, error) {
	panic("implement me")
}

func (s spaceService) GetUserSpaces(ctx context.Context, userID string) ([]models.Space, error) {
	panic("implement me")
}
