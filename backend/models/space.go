package models

import "go.mongodb.org/mongo-driver/v2/bson"

type UserIdWithRole struct {
	UserID bson.ObjectID `bson:"userId" json:"userId"`
	Role   string        `bson:"role" json:"role"`
}

type Space struct {
	BaseEntity  `bson:",inline" json:",inline"`
	Name        string           `bson:"name" json:"name"`
	Description string           `bson:"description" json:"description"`
	OwnerID     bson.ObjectID    `bson:"ownerId" json:"ownerId"`
	Pages       []bson.ObjectID  `bson:"pages,omitempty" json:"pages,omitempty"`
	Shared      []UserIdWithRole `bson:"shared"`
}
