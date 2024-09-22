package repositories

import (
	"backend/models"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type PageRepository struct {
	*GenericRepository[*models.Page]
}

func NewPageRepository(collection *mongo.Collection) *PageRepository {
	repo := &PageRepository{
		GenericRepository: NewGenericRepository[*models.Page](collection),
	}

	err := repo.createIndexes()
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	return repo
}

func (r *PageRepository) FindBySpaceId(ctx context.Context, spaceID bson.ObjectID) ([]models.Page, error) {
	// Ruft die Find-Methode auf, die Pointer zurückgibt
	pointerPages, err := r.Find(ctx, bson.M{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	// Wandelt die Liste von *models.Page zu []models.Page um
	pages := make([]models.Page, len(pointerPages))
	for i, p := range pointerPages {
		pages[i] = *p // Dereferenzierung des Pointers
	}

	return pages, nil
}

func (r *PageRepository) BuildPageTree(pages []*models.Page) []*models.Page {
	// Map pages by ParentPageID
	pageMap := make(map[bson.ObjectID][]*models.Page)
	var rootPages []*models.Page

	for _, page := range pages {
		if page.ParentPageID != nil {
			pageMap[*page.ParentPageID] = append(pageMap[*page.ParentPageID], page)
		} else {
			rootPages = append(rootPages, page)
		}
	}

	// Recursively build the tree
	var buildSubPages func(*models.Page)
	buildSubPages = func(parent *models.Page) {
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
	_, err := r.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return err
	}
	log.Println("Index created successfully for spaceId")
	return nil
}
