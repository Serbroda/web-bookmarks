package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*model.Space]
	pageRepo *PageRepository
}

func NewSpaceRepository(collection *mongo.Collection, pageRepo *PageRepository) *SpaceRepository {
	return &SpaceRepository{
		GenericRepository: NewGenericRepository[*model.Space](collection),
		pageRepo:          pageRepo,
	}
}

// FindByIdWithPages nutzt die allgemeine Find-Methode, um die Pages zu laden
func (r *SpaceRepository) FindByIdWithPages(ctx context.Context, id bson.ObjectID) (*model.Space, error) {
	// Space durch ID finden
	var space *model.Space
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&space)
	if err != nil {
		return nil, err
	}

	// Pages für den Space mit dynamischem Filter laden
	pages, err := r.pageRepo.Find(ctx, bson.M{"spaceId": space.ID})
	if err != nil {
		return nil, err
	}

	// Pages IDs dem Space hinzufügen
	space.Pages = make([]bson.ObjectID, len(pages))
	for i, page := range pages {
		space.Pages[i] = page.ID
	}

	return space, nil
}
