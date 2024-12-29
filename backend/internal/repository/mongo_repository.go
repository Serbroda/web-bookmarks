package repository

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepository struct {
	client             *mongo.Client
	db                 *mongo.Database
	bookmarkCollection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) (Repository, error) {
	return &MongoRepository{
		bookmarkCollection: database.Collection("bookmarks"),
	}, nil
}

// Implementation
func (m *MongoRepository) CreateBookmark(ctx context.Context, bookmark models.Bookmark) (models.Bookmark, error) {
	if bookmark.ID == "" {
		bookmark.ID = bson.NewObjectID().Hex()
	}

	doc := bson.M{
		"_id":         bookmark.ID,
		"url":         bookmark.URL,
		"title":       bookmark.Title,
		"description": bookmark.Description,
	}

	_, err := m.bookmarkCollection.InsertOne(ctx, doc)
	return bookmark, err
}

func (m *MongoRepository) GetAllBookmarks(ctx context.Context) ([]models.Bookmark, error) {
	cursor, err := m.bookmarkCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var results []models.Bookmark
	for cursor.Next(ctx) {
		var doc struct {
			ID          string `bson:"_id"`
			URL         string `bson:"url"`
			Title       string `bson:"title"`
			Description string `bson:"description"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, models.Bookmark{
			ID:          doc.ID,
			URL:         doc.URL,
			Title:       doc.Title,
			Description: doc.Description,
		})
	}
	return results, cursor.Err()
}

func (m *MongoRepository) GetBookmarkByID(ctx context.Context, id string) (models.Bookmark, error) {
	var doc struct {
		ID          string `bson:"_id"`
		URL         string `bson:"url"`
		Title       string `bson:"title"`
		Description string `bson:"description"`
	}
	err := m.bookmarkCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	if err != nil {
		return models.Bookmark{}, err
	}
	return models.Bookmark{
		ID:          doc.ID,
		URL:         doc.URL,
		Title:       doc.Title,
		Description: doc.Description,
	}, nil
}
