package models

type Space struct {
	Base
	OwnerId     uint       `json:"ownerId"`
	Owner       User       `json:"-"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
	Groups      []Group    `json:"groups"`
}
