package services

import (
	"fmt"

	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/models"
)

func FindLinksByGroupId(groupId string) []models.Link {
	var links []models.Link
	database.GetConnection().Where("group_id = ?", groupId).Find(&links)
	return links
}

func FindLinks(ownderId uint, order string, limit int) []models.Link {
	var links []models.Link
	sql := `select links.* 
		from links 
			inner join groups on groups.id = links.group_id 
		where links.deleted_at is null and groups.owner_id = ?`

	if order != "" {
		sql += " ORDER BY " + order + " "
	}
	if limit > 0 {
		sql += fmt.Sprint(" LIMIT ", limit)
	}
	database.GetConnection().Raw(sql, ownderId).Scan(&links)
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
