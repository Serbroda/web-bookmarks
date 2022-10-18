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

func FindGroupsByOwnerId(ownerId uint, order string, limit int) []models.Group {
	var entities []models.Group
	if order == "" {
		order = "id"
	}
	database.GetConnection().Where("owner_id = ?", ownerId).Order(order).Limit(limit).Find(&entities)
	return entities
}
