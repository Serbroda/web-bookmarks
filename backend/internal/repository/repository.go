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
	Save(ctx context.Context, entity *T) (string, error)
	SaveWithId(ctx context.Context, entity *T) (string, error)
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

func (r *MongoRepository[T]) Save(ctx context.Context, entity *T) (string, error) {
	return r.SaveWithId(ctx, entity, "_id")
}

func (r *MongoRepository[T]) SaveWithId(ctx context.Context, entity *T, idField string) (string, error) {
	// Marshal the entity into a BSON document
	doc := bson.M{}
	data, _ := bson.Marshal(entity)
	err := bson.Unmarshal(data, &doc)
	if err != nil {
		return "", err
	}

	now := time.Now()
	doc["updatedAt"] = now

	// Überprüfen, ob eine ID vorhanden ist
	if id, ok := doc[idField]; ok {
		// Update durch Upsert (falls das Dokument existiert, wird es aktualisiert)
		filter := bson.M{idField: id}
		update := bson.M{"$set": doc}
		opts := options.Update().SetUpsert(true)
		res, err := r.collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Failed to upsert document: %v", err)
			return "", err
		}

		// Überprüfen, ob ein neues Dokument eingefügt wurde (Upsert)
		if res.UpsertedID != nil {
			upsertedID := res.UpsertedID.(bson.ObjectID).Hex()
			log.Println("Document upserted successfully with ID:", upsertedID)
			return upsertedID, nil
		}

		log.Println("Document updated successfully")
		return id.(bson.ObjectID).Hex(), nil

	} else {
		// Neue Dokumente einfügen
		doc["createdAt"] = now

		res, err := r.collection.InsertOne(ctx, doc)
		if err != nil {
			log.Printf("Failed to insert document: %v", err)
			return "", err
		}

		// Gebe die InsertedID als Hex-String zurück
		insertedID := res.InsertedID.(bson.ObjectID).Hex()
		log.Println("Document inserted successfully with ID:", insertedID)
		return insertedID, nil
	}
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
