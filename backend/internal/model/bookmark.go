package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Bookmark struct {
	BaseEntity  `bson:",inline" json:",inline"`
	Title       string        `bson:"title" json:"title"`
	URL         string        `bson:"url" json:"url"`
	Description string        `bson:"description,omitempty" json:"description,omitempty"`
	PageId      bson.ObjectID `bson:"pageId" json:"pageId"` // Verweis auf Kategorie
	Tags        []string      `bson:"tags,omitempty" json:"tags,omitempty"`
}
