package services

import (
	"errors"
	"webcrate/database"
	"webcrate/models"
)

var ErrEntityNotFound = errors.New("can not find entity")

func FindGroupById(id string) (models.Group, error) {
	var entity models.Group
	result := database.GetConnection().Where("id = ?", id).Find(&entity)

	if result.RowsAffected == 0 {
		return models.Group{}, ErrEntityNotFound
	}
	return entity, nil
}

func FindGroupsByOwnerId(ownerId uint) []models.Group {
	var entities []models.Group
	database.GetConnection().Where("owner_id = ?", ownerId).Find(&entities)
	return entities
}
