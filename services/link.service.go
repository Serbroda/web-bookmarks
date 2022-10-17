package services

import (
	"webcrate/database"
	"webcrate/models"
)

func FindLinksByGroupId(groupId string) []models.Link {
	var links []models.Link
	database.GetConnection().Where("group_id = ?", groupId).Find(&links)
	return links
}

func FindLinkById(id string) (models.Link, error) {
	var entity models.Link
	result := database.GetConnection().Where("id = ?", id).Find(&entity)

	if result.RowsAffected == 0 {
		return models.Link{}, ErrEntityNotFound
	}
	return entity, nil
}
