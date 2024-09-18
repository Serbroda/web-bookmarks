package repository

import (
	"backend/internal/events"
	"backend/internal/model"
	"backend/internal/util"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpaceRepository struct {
	*GenericRepository[*model.Space]
}

func NewSpaceRepository(collection *mongo.Collection, dispatcher *events.EventDispatcher) *SpaceRepository {
	repo := &SpaceRepository{
		GenericRepository: NewGenericRepository[*model.Space](collection, dispatcher, "Space"),
	}

	dispatcher.RegisterListener("PageInsert", func(event events.Event) {
		page := event.Data.(*model.Page)

		space, err := repo.FindByID(context.TODO(), page.SpaceID)
		if err != nil {
			return
		}
		space.Pages = append(space.Pages, page.ID)
		err = repo.Save(context.TODO(), space)
		if err != nil {
			return
		}
	})

	dispatcher.RegisterListener("PageUpdate", func(event events.Event) {
		page := event.Data.(*model.Page)
		old := event.OldData.(*model.Page)

		if old.SpaceID != page.SpaceID {
			space, err := repo.FindByID(context.TODO(), page.SpaceID)
			if err != nil {
				return
			}
			space.Pages = util.Remove(space.Pages, func(item bson.ObjectID) bool {
				return item == page.ID
			})
			err = repo.Save(context.TODO(), space)
			if err != nil {
				return
			}
		}
	})

	return repo
}
