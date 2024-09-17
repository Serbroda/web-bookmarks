package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type FeatureService struct {
	spaceRepo *repository.SpaceRepository
	pageRepo  *repository.PageRepository
}

func NewFeatureService(spaceRepo *repository.SpaceRepository, pageRepo *repository.PageRepository) *FeatureService {
	return &FeatureService{
		spaceRepo: spaceRepo,
		pageRepo:  pageRepo,
	}
}

func (s *FeatureService) GetSpaceById(ctx context.Context, id bson.ObjectID) (model.Space, error) {
	space, err := s.spaceRepo.FindByID(ctx, id)
	if err != nil {
		return model.Space{}, err
	}

	pages, err := s.pageRepo.FindBySpaceId(ctx, space.ID)
	if err != nil {
		return model.Space{}, err
	}

	space.Pages = make([]bson.ObjectID, len(pages))
	for i, page := range pages {
		space.Pages[i] = page.ID
	}
	return *space, nil
}
