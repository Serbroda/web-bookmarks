package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository interface {
	GenericRepository[model.Space] // Inklusive der CRUD-Methoden
	FindBySpaceName(ctx context.Context, name string) (*model.Space, error)
}

type MongoSpaceRepository struct {
	*MongoRepository[model.Space] // Embedding des generischen Repositories
	collection                    *mongo.Collection
}

func NewMongoSpaceRepository(collection *mongo.Collection) *MongoSpaceRepository {
	return &MongoSpaceRepository{
		MongoRepository: NewMongoRepository[model.Space](collection), // Initialisiere das generische Repository
		collection:      collection,
	}
}

func (r *MongoSpaceRepository) FindBySpaceName(ctx context.Context, name string) (*model.Space, error) {
	var result model.Space
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
