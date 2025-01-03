package repository

import (
	"context"
	"fmt"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"time"
)

type MongoRepository struct {
	client    *mongo.Client
	db        *mongo.Database
	collBmk   *mongo.Collection
	collSpace *mongo.Collection
}

func NewMongoRepository(uri string, dbName string) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	db := client.Database(dbName)

	return &MongoRepository{
		client:    client,
		db:        db,
		collBmk:   db.Collection("bookmarks"),
		collSpace: db.Collection("spaces"),
	}, nil
}

func (m *MongoRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	err := save[*models.User](ctx, m.collSpace, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (m *MongoRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoRepository) CreateSpace(ctx context.Context, space models.Space) (models.Space, error) {
	err := save[*models.Space](ctx, m.collSpace, &space)
	if err != nil {
		return models.Space{}, err
	}
	return space, nil
}

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

	_, err := m.collBmk.InsertOne(ctx, doc)
	return bookmark, err
}

func (m *MongoRepository) GetAllBookmarks(ctx context.Context) ([]models.Bookmark, error) {
	cursor, err := m.collBmk.Find(ctx, bson.M{})
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
	err := m.collBmk.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
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

func save[T models.BaseEntityInterface](ctx context.Context, collection *mongo.Collection, entity T) error {
	now := time.Now()
	entity.SetUpdatedAt(now)

	if entity.GetID() != "" {
		entity.SetCreatedAt(now)
		entity.SetID(bson.NewObjectID().Hex()) // Setze eine neue ID
	}

	filter := bson.M{"_id": entity.GetID()}
	update := bson.M{"$set": entity}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}
