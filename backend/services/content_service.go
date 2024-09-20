package services

import (
	"backend/models"
	"backend/repositories"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ContentService struct {
	spaceRepo    *repositories.SpaceRepository
	pageRepo     *repositories.PageRepository
	bookmarkRepo *repositories.BookmarkRepository
}

func NewContentService(
	spaceRepo *repositories.SpaceRepository,
	pageRepo *repositories.PageRepository,
	bookmarkRepo *repositories.BookmarkRepository) *ContentService {
	return &ContentService{
		spaceRepo:    spaceRepo,
		pageRepo:     pageRepo,
		bookmarkRepo: bookmarkRepo,
	}
}

func (s *ContentService) CreateSpace(ctx context.Context, space *models.Space) error {
	return s.spaceRepo.Save(ctx, space)
}

func (s *ContentService) GetSpaceById(ctx context.Context, id bson.ObjectID) (models.Space, error) {
	space, err := s.spaceRepo.FindByID(ctx, id)
	if err != nil {
		return models.Space{}, err
	}

	pages, err := s.pageRepo.FindBySpaceId(ctx, space.ID)
	if err != nil {
		return models.Space{}, err
	}

	space.Pages = make([]bson.ObjectID, len(pages))
	for i, page := range pages {
		space.Pages[i] = page.ID
	}
	return *space, nil
}

func (s *ContentService) GetSpacesForUser(userId bson.ObjectID) ([]*models.Space, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"ownerId": userId},
			{"shared.userId": userId},
		},
	}

	founds, err := s.spaceRepo.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return founds, nil
}

func (s *ContentService) DeleteSpace(ctx context.Context, id bson.ObjectID) error {
	return s.spaceRepo.Delete(ctx, id)
}
