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
	Name        string `json:"name" xml:"name" form:"name"`
	Url         string `json:"url" xml:"url" form:"url"`
	Description string `json:"description" xml:"description" form:"description"`
	GroupId     string `json:"groupId" xml:"groupId" form:"groupId"`
}
