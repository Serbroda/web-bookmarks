package models

import (
	"database/sql"
)

type Space struct {
	BaseModel
	ShortID     string         `db:"short_id"`
	OwnerID     int64          `db:"owner_id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Visibility  string         `db:"visibility"`
}
