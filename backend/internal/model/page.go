package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Page struct {
	BaseEntity   `bson:",inline" json:",inline"`
	Name         string          `bson:"name" json:"name"`
	SpaceID      bson.ObjectID   `bson:"spaceId" json:"spaceId"`
	ParentPageID *bson.ObjectID  `bson:"parentPageId,omitempty" json:"parentPageId,omitempty"`
	SubPages     []bson.ObjectID `bson:"subPages,omitempty" json:"subPages,omitempty"`
	Bookmarks    []bson.ObjectID `bson:"bookmarks,omitempty" json:"bookmarks,omitempty"`
}
