package mongodb

import (
	"backend/internal"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericMongoRepository[*internal.Space]
}

func NewSpaceRepository(collection *mongo.Collection) *SpaceRepository {
	repo := &SpaceRepository{
		GenericMongoRepository: NewGenericRepository[*internal.Space](collection),
	}
	return repo
}
