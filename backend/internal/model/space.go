package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Space struct {
	BaseEntity  `bson:",inline" json:",inline"`
	Name        string          `bson:"name" json:"name"`
	Description string          `bson:"description" json:"description"`
	OwnerID     bson.ObjectID   `bson:"ownerId" json:"ownerId"`                           // User who owns the Space
	SharedWith  []bson.ObjectID `bson:"sharedWith,omitempty" json:"sharedWith,omitempty"` // Users with whom the space is shared
	Categories  []bson.ObjectID `bson:"categories,omitempty" json:"categories,omitempty"` // List of Category IDs
}
