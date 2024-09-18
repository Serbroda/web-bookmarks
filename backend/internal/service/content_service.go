package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ContentService struct {
	spaceRepo    *repository.SpaceRepository
	pageRepo     *repository.PageRepository
	bookmarkRepo *repository.BookmarkRepository
}

func NewFeatureService(spaceRepo *repository.SpaceRepository, pageRepo *repository.PageRepository, bookmarkRepo *repository.BookmarkRepository) *ContentService {
	return &ContentService{
		spaceRepo:    spaceRepo,
		pageRepo:     pageRepo,
		bookmarkRepo: bookmarkRepo,
	}
}

func (s *ContentService) GetSpaceById(ctx context.Context, id bson.ObjectID) (model.Space, error) {
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
