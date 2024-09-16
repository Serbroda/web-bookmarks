package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type GenericRepository[T any] interface {
	Insert(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
}

type MongoRepository[T any] struct {
	collection *mongo.Collection
}

func NewMongoRepository[T any](collection *mongo.Collection) *MongoRepository[T] {
	return &MongoRepository[T]{collection: collection}
}

func (r *MongoRepository[T]) Insert(ctx context.Context, entity *T) error {
	_, err := r.collection.InsertOne(ctx, entity)
	return err
}

func (r *MongoRepository[T]) FindByID(ctx context.Context, id string) (*T, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result T
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MongoRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var results []T
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MongoRepository[T]) Update(ctx context.Context, id string, entity *T) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": objectID}, entity)
	return err
}

func (r *MongoRepository[T]) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
