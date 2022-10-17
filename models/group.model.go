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
	Icon        string     `json:"icon"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	Links       []Link     `json:"links"`
}

type CreateGroupDto struct {
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
