package mongodb

import (
	"backend/internal"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type PageRepository struct {
	*GenericMongoRepository[*internal.Page]
}

func NewPageRepository(collection *mongo.Collection) *PageRepository {
	repo := &PageRepository{
		GenericMongoRepository: NewGenericRepository[*internal.Page](collection),
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *PageRepository) FindBySpaceId(ctx context.Context, spaceID bson.ObjectID) ([]internal.Page, error) {
	// Ruft die Find-Methode auf, die Pointer zurückgibt
	pointerPages, err := r.Find(ctx, bson.M{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	// Wandelt die Liste von *models.Page zu []models.Page um
	pages := make([]internal.Page, len(pointerPages))
	for i, p := range pointerPages {
		pages[i] = *p // Dereferenzierung des Pointers
	}

	return pages, nil
}

func (r *PageRepository) BuildPageTree(pages []*internal.Page) []*internal.Page {
	// Map pages by ParentPageID
	pageMap := make(map[bson.ObjectID][]*internal.Page)
	var rootPages []*internal.Page

	for _, page := range pages {
		if page.ParentPageID != nil {
			pageMap[*page.ParentPageID] = append(pageMap[*page.ParentPageID], page)
		} else {
			rootPages = append(rootPages, page)
		}
	}

	// Recursively build the tree
	var buildSubPages func(*internal.Page)
	buildSubPages = func(parent *internal.Page) {
		if subPages, ok := pageMap[parent.ID]; ok {
			parent.SubPages = subPages
			for _, subPage := range subPages {
				buildSubPages(subPage)
			}
		}
	}

	for _, root := range rootPages {
		buildSubPages(root)
	}

	return rootPages
}

func (r *PageRepository) createIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"spaceId": 1}, // 1 für aufsteigender Index
		Options: options.Index().SetName("idx_pages_spaceId"),
	}
	_, err := r.Collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for spaceId")
	return nil
}
