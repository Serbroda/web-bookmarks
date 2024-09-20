package repositories

import (
	"backend/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*models.Space]
}

func NewSpaceRepository(collection *mongo.Collection) *SpaceRepository {
	repo := &SpaceRepository{
		GenericRepository: NewGenericRepository[*models.Space](collection),
	}
	return repo
}
