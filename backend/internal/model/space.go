package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Space struct {
	BaseEntity  `bson:",inline" json:",inline"`
	Name        string          `bson:"name" json:"name"`
	Description string          `bson:"description" json:"description"`
	OwnerID     bson.ObjectID   `bson:"ownerId" json:"ownerId"`
	Pages       []bson.ObjectID `bson:"pages,omitempty" json:"pages,omitempty"`
	SharedWith  []bson.ObjectID `bson:"sharedWith,omitempty" json:"sharedWith,omitempty"`
}
