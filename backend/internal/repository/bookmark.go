package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type BookmarkRepository interface {
	GenericRepository[model.Bookmark] // Inklusive der CRUD-Methoden
	FindByBookmarkName(ctx context.Context, name string) (*model.Bookmark, error)
}

type MongoBookmarkRepository struct {
	*MongoRepository[model.Bookmark] // Embedding des generischen Repositories
	collection                       *mongo.Collection
}

func NewMongoBookmarkRepository(collection *mongo.Collection) *MongoBookmarkRepository {
	repo := &MongoBookmarkRepository{
		MongoRepository: NewMongoRepository[model.Bookmark](collection),
		collection:      collection,
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *MongoBookmarkRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"pageId": 1},
		Options: options.Index().SetName("idx_pageId"),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for pageId")
	return nil
}
