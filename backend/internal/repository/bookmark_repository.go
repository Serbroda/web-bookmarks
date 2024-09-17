package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type BookmarkRepository struct {
	*GenericRepository[*model.Bookmark]
}

func NewBookmarkRepository(collection *mongo.Collection) *BookmarkRepository {
	repo := &BookmarkRepository{
		GenericRepository: NewGenericRepository[*model.Bookmark](collection),
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *BookmarkRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"pageId": 1}, // 1 für aufsteigender Index
		Options: options.Index().SetName("idx_bookmarks_pageId"),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for pageId")
	return nil
}
