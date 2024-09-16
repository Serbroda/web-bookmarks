package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type PageRepository interface {
	GenericRepository[model.Page] // Inklusive der CRUD-Methoden
	FindByPageName(ctx context.Context, name string) (*model.Page, error)
}

type MongoPageRepository struct {
	*MongoRepository[model.Page] // Embedding des generischen Repositories
	collection                   *mongo.Collection
}

func NewMongoPageRepository(collection *mongo.Collection) *MongoPageRepository {
	repo := &MongoPageRepository{
		MongoRepository: NewMongoRepository[model.Page](collection),
		collection:      collection,
	}

	// Erstelle den Index für spaceId bei der Initialisierung
	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *MongoPageRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"spaceId": 1}, // 1 für aufsteigender Index
		Options: options.Index().SetName("idx_spaceId"),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for spaceId")
	return nil
}
