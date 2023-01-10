package models

type Role struct {
	BaseModel
	Name        string  `db:"name"`
	Description *string `db:"description"`
}
