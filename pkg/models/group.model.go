package models

type Visibility string

const (
	Private Visibility = "private"
	Public             = "public"
)

type Group struct {
	Base
	OwnerId     uint       `json:"ownerId"`
	Owner       User       `json:"-"`
	SpaceId     uint       `json:"spaceId"`
	Space       Space      `json:"-"`
	Icon        string     `json:"icon"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	Links       []Link     `json:"links"`
}

type CreateGroupDto struct {
	Icon        string `json:"icon" xml:"icon" form:"icon"`
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
}

type ChangeGroupVisibility struct {
	Visibility Visibility `json:"visibility" xml:"visibility" form:"visibility"`
}
