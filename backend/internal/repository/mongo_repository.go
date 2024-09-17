package repository

import (
	"backend/internal/events"
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type GenericRepository[T model.BaseEntityInterface] struct {
	collection *mongo.Collection
	dispatcher *events.EventDispatcher
	saveEvent  string
}

func NewGenericRepository[T model.BaseEntityInterface](collection *mongo.Collection, dispatcher *events.EventDispatcher, saveEvent string) *GenericRepository[T] {
	return &GenericRepository[T]{
		collection: collection,
		dispatcher: dispatcher,
		saveEvent:  saveEvent,
	}
}

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

	r.dispatcher.Dispatch(events.Event{
		Name: r.saveEvent,
		Data: entity,
	})
	return err
}

func (r *GenericRepository[T]) FindByID(ctx context.Context, id bson.ObjectID) (T, error) {
	var entity T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	return entity, err
}

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

func (r *GenericRepository[T]) Delete(ctx context.Context, id bson.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *GenericRepository[T]) Find(ctx context.Context, filter interface{}) ([]T, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var entities []T
	if err = cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	return entities, nil
}
