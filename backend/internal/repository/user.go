package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type UserRepository interface {
	GenericRepository[model.User] // Inklusive der CRUD-Methoden
	FindByUserName(ctx context.Context, name string) (*model.User, error)
}

type MongoUserRepository struct {
	*MongoRepository[model.User] // Embedding des generischen Repositories
	collection                   *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	repo := &MongoUserRepository{
		MongoRepository: NewMongoRepository[model.User](collection),
		collection:      collection,
	}

	// Erstelle den Index für spaceId bei der Initialisierung
	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *MongoUserRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"username": 1}, // 1 für aufsteigender Index
		Options: options.Index().
			SetName("uc_users_username").
			SetUnique(true),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for spaceId")
	return nil
}
