package internal

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type BaseEntityInterface interface {
	SetID(id bson.ObjectID)
	GetID() bson.ObjectID
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

type BaseEntity struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"createdAt" json:"-"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"-"`
}

func (e *BaseEntity) SetID(id bson.ObjectID) {
	e.ID = id
}

func (e *BaseEntity) GetID() bson.ObjectID {
	return e.ID
}

func (e *BaseEntity) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

func (e *BaseEntity) SetUpdatedAt(t time.Time) {
	e.UpdatedAt = t
}
