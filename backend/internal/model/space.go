package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Space struct {
	BaseEntity `bson:",inline"` // Inline macht die Felder von BaseEntity direkt verfügbar
	Name       string           `bson:"name" json:"name"`
	OwnerID    bson.ObjectID    `bson:"ownerId" json:"ownerId"`
}
