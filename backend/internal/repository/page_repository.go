package repository

import (
	"backend/internal/events"
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type PageRepository struct {
	*GenericRepository[*model.Page]
}

func NewPageRepository(collection *mongo.Collection, dispatcher *events.EventDispatcher) *PageRepository {
	repo := &PageRepository{
		GenericRepository: NewGenericRepository[*model.Page](collection, dispatcher, "PageSaved"),
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *PageRepository) FindBySpaceId(ctx context.Context, spaceID bson.ObjectID) ([]model.Page, error) {
	// Ruft die Find-Methode auf, die Pointer zurückgibt
	pointerPages, err := r.Find(ctx, bson.M{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	// Wandelt die Liste von *model.Page zu []model.Page um
	pages := make([]model.Page, len(pointerPages))
	for i, p := range pointerPages {
		pages[i] = *p // Dereferenzierung des Pointers
	}

	return pages, nil
}

func (r *PageRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"spaceId": 1}, // 1 für aufsteigender Index
		Options: options.Index().SetName("idx_pages_spaceId"),
	}
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for spaceId")
	return nil
}
