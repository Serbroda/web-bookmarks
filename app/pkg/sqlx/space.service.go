package sqlx

import (
	"github.com/Serbroda/ragbag/pkg/model"
	"github.com/jmoiron/sqlx"
)

type SpaceServiceSqlx struct {
	DB *sqlx.DB
}

func (s SpaceServiceSqlx) FindOne(id int64) (*model.Space, error) {
	var entity model.Space
	err := s.DB.Get(&entity, "SELECT * FROM spaces WHERE id = $1 AND deleted_at is null LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (s SpaceServiceSqlx) FindOneByShortId(shortId string) (*model.Space, error) {
	var entity model.Space
	err := s.DB.Get(&entity, "SELECT * FROM spaces WHERE short_id = $1 AND deleted_at is null LIMIT 1", shortId)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (s SpaceServiceSqlx) Create(owner int64, params model.CreateSpace) (*model.Space, error) {
	//TODO implement me
	panic("implement me")
}
