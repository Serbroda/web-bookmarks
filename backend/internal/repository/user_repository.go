package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type UserRepository struct {
	*GenericRepository[*model.User]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	repo := &UserRepository{
		GenericRepository: NewGenericRepository[*model.User](collection),
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *UserRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"email": 1}, // 1 für aufsteigender Index
		Options: options.Index().
			SetName("uc_users_email").
			SetUnique(true),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for email")

	indexModel = mongo.IndexModel{
		Keys: bson.M{"username": 1}, // 1 für aufsteigender Index
		Options: options.Index().
			SetName("uc_users_username").
			SetUnique(true),
	}
	_, err = r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for username")
	return nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	found, err := r.FindOne(ctx, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	return *found, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	found, err := r.FindOne(ctx, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	return *found, err
}
