package models

import "go.mongodb.org/mongo-driver/v2/bson"

type SpaceVisibility = string

const (
	SpaceVisibilityPublic  SpaceVisibility = "PUBLIC"
	SpaceVisibilityPrivate SpaceVisibility = "PRIVATE"
)

/*type UserIdWithRole struct {
	UserID bson.ObjectID `bson:"userId" json:"userId"`
	Role   string        `bson:"role" json:"role"`
}*/

type Space struct {
	BaseEntity  `bson:",inline" json:",inline"`
	Name        string          `bson:"name" json:"name"`
	Description *string         `bson:"description" json:"description"`
	Visibility  SpaceVisibility `bson:"visibility" json:"visibility"`
	OwnerID     bson.ObjectID   `bson:"ownerId" json:"ownerId"`
	//Shared      []UserIdWithRole `bson:"shared"`
}
