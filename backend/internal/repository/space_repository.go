package repository

import (
	"backend/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*model.Space]
}

func NewSpaceRepository(collection *mongo.Collection) *SpaceRepository {
	repo := &SpaceRepository{
		GenericRepository: NewGenericRepository[*model.Space](collection),
	}
	return repo
}
