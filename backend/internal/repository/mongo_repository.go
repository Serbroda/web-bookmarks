package repository

import (
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

// GenericRepository verwendet *T als generischen Typ
type GenericRepository[T model.BaseEntityInterface] struct {
	collection *mongo.Collection
}

// NewGenericRepository erstellt ein neues generisches Repository
func NewGenericRepository[T model.BaseEntityInterface](collection *mongo.Collection) *GenericRepository[T] {
	return &GenericRepository[T]{collection: collection}
}

// Save führt entweder ein Insert oder ein Update durch
func (r *GenericRepository[T]) Save(ctx context.Context, entity T) error {
	now := time.Now()
	entity.SetUpdatedAt(now)

	if entity.GetID().IsZero() {
		entity.SetCreatedAt(now)
		entity.SetID(bson.NewObjectID()) // Setze eine neue ID
	}

	filter := bson.M{"_id": entity.GetID()}
	update := bson.M{"$set": entity}
	opts := options.Update().SetUpsert(true)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

// FindByID findet ein Dokument anhand der ID
func (r *GenericRepository[T]) FindByID(ctx context.Context, id bson.ObjectID) (T, error) {
	var entity T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	return entity, err
}

// FindAll gibt alle Dokumente in der Sammlung zurück
func (r *GenericRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var entities []T
	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}

// Delete löscht ein Dokument anhand der ID
func (r *GenericRepository[T]) Delete(ctx context.Context, id bson.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
