package repository

import (
	"backend/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*model.Space]
}

// NewSpaceRepository erstellt ein neues SpaceRepository
func NewSpaceRepository(collection *mongo.Collection) *SpaceRepository {
	return &SpaceRepository{
		GenericRepository: NewGenericRepository[*model.Space](collection),
	}
}
