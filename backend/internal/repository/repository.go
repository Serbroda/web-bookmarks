package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

type GenericRepository[T any] interface {
	Save(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	FindAll(ctx context.Context) ([]T, error)
	Delete(ctx context.Context, id string) error
}

type MongoRepository[T any] struct {
	collection *mongo.Collection
}

func NewMongoRepository[T any](collection *mongo.Collection) *MongoRepository[T] {
	return &MongoRepository[T]{collection: collection}
}

func (r *MongoRepository[T]) Save(ctx context.Context, entity *T) error {
	return r.SaveWithId(ctx, entity, "_id")
}

func (r *MongoRepository[T]) SaveWithId(ctx context.Context, entity *T, idField string) error {
	// Reflexion, um die _id des Objekts zu ermitteln
	doc := bson.M{}
	data, _ := bson.Marshal(entity)
	err := bson.Unmarshal(data, &doc)
	if err != nil {
		return err
	}

	now := time.Now()
	doc["updatedAt"] = now

	// Überprüfen, ob eine ID vorhanden ist
	if id, ok := doc[idField]; ok {
		// Update durch Upsert (falls das Dokument existiert, wird es aktualisiert)
		filter := bson.M{idField: id}
		update := bson.M{"$set": doc}
		opts := options.Update().SetUpsert(true)
		_, err := r.collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Failed to upsert document: %v", err)
			return err
		}
		log.Println("Document updated successfully")
	} else {
		doc["createdAt"] = now

		_, err := r.collection.InsertOne(ctx, doc)
		if err != nil {
			log.Printf("Failed to insert document: %v", err)
			return err
		}
		log.Println("Document inserted successfully")
	}
	return nil
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

func (r *MongoRepository[T]) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
