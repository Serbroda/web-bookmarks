package models

import (
	"time"
)

type BaseEntityInterface interface {
	SetID(id string)
	GetID() string
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type BaseEntity struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time `bson:"createdAt" json:"-"`
	UpdatedAt time.Time `bson:"updatedAt" json:"-"`
}

func (e *BaseEntity) SetID(id string) {
	e.ID = id
}

func (e *BaseEntity) GetID() string {
	return e.ID
}

func (e *BaseEntity) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

func (e *BaseEntity) SetUpdatedAt(t time.Time) {
	e.UpdatedAt = t
}
