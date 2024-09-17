package repository

import (
	"backend/internal/events"
	"backend/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*model.Space]
}

func NewSpaceRepository(collection *mongo.Collection, dispatcher *events.EventDispatcher) *SpaceRepository {
	repo := &SpaceRepository{
		GenericRepository: NewGenericRepository[*model.Space](collection, dispatcher, "SpaceSaved"),
	}

	dispatcher.RegisterListener("PageSaved", func(event events.Event) {
		page := event.Data.(*model.Page) // Die Page aus dem Event extrahieren

		space, err := repo.FindByID(context.TODO(), page.SpaceID)
		if err != nil {
			return
		}
		space.Pages = append(space.Pages, page.ID)
	})

	return repo
}
