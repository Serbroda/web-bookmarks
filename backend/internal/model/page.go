package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Page struct {
	BaseEntity       `bson:",inline" json:",inline"`
	Name             string          `bson:"name" json:"name"`
	SpaceID          bson.ObjectID   `bson:"spaceId" json:"spaceId"`                                       // Verweis auf Space
	ParentCategoryID *bson.ObjectID  `bson:"parentCategoryId,omitempty" json:"parentCategoryId,omitempty"` // Verweis auf übergeordnete Kategorie (falls vorhanden)
	Subcategories    []bson.ObjectID `bson:"subcategories,omitempty" json:"subcategories,omitempty"`       // List von Subkategorien
	Bookmarks        []bson.ObjectID `bson:"bookmarks,omitempty" json:"bookmarks,omitempty"`               // List von Bookmark-IDs
}
