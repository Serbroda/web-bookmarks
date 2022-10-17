package models

type Link struct {
	Base
	GroupId     string     `json:"groupId"`
	Group       Group      `json:"-"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	Description string     `json:"description"`
	Visibility  Visibility `json:"visibility"`
}

type CreateLinkDto struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	GroupId     string `json:"groupId"`
}
