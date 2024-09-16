package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type BaseEntity struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}
