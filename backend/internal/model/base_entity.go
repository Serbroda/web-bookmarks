package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// BaseEntityInterface stellt die Methode SetID bereit
type BaseEntityInterface interface {
	SetID(id bson.ObjectID)
	GetID() bson.ObjectID
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
}

// BaseEntity enthält die allgemeinen Felder und Methoden
type BaseEntity struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"createdAt" json:"-"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"-"`
}

// SetID implementiert BaseEntityInterface.SetID
func (e *BaseEntity) SetID(id bson.ObjectID) {
	e.ID = id
}

// GetID gibt die aktuelle ID zurück
func (e *BaseEntity) GetID() bson.ObjectID {
	return e.ID
}

// SetCreatedAt setzt den CreatedAt-Timestamp
func (e *BaseEntity) SetCreatedAt(t time.Time) {
	e.CreatedAt = t
}

// SetUpdatedAt setzt den UpdatedAt-Timestamp
func (e *BaseEntity) SetUpdatedAt(t time.Time) {
	e.UpdatedAt = t
}
