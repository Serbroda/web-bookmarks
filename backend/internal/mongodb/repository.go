package mongodb

import (
	"backend/internal"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

var (
	ErrMoreThanOneEntity = errors.New("more than one entity found")
)

type GenericMongoRepository[T internal.BaseEntityInterface] struct {
	Collection *mongo.Collection
}

func NewGenericRepository[T internal.BaseEntityInterface](collection *mongo.Collection) *GenericMongoRepository[T] {
	return &GenericMongoRepository[T]{
		Collection: collection,
	}
}

func (r *GenericMongoRepository[T]) Save(ctx context.Context, entity T) error {
	now := time.Now()
	entity.SetUpdatedAt(now)

	if entity.GetID().IsZero() {
		entity.SetCreatedAt(now)
		entity.SetID(bson.NewObjectID()) // Setze eine neue ID
	}

	filter := bson.M{"_id": entity.GetID()}
	update := bson.M{"$set": entity}
	opts := options.Update().SetUpsert(true)

	_, err := r.Collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *GenericMongoRepository[T]) FindByID(ctx context.Context, id bson.ObjectID) (T, error) {
	var entity T
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	return entity, err
}

func (r *GenericMongoRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var entities []T
	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *GenericMongoRepository[T]) Find(ctx context.Context, filter interface{}) ([]T, error) {
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var entities []T
	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *GenericMongoRepository[T]) FindOne(ctx context.Context, filter interface{}) (*T, error) {
	entities, err := r.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	size := len(entities)
	if size == 0 {
		return nil, mongo.ErrNoDocuments
	} else if size > 1 {
		return nil, ErrMoreThanOneEntity
	}
	return &entities[0], nil
}

func (r *GenericMongoRepository[T]) Delete(ctx context.Context, id bson.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
