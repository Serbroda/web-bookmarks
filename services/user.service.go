package services

import (
	"webcrate/database"
	"webcrate/models"
)

func FindUserById(id uint) (models.User, error) {
	var entity models.User
	result := database.GetConnection().Where("id = ?", id).Find(&entity)

	if result.RowsAffected == 0 {
		return models.User{}, ErrEntityNotFound
	}
	return entity, nil
}
