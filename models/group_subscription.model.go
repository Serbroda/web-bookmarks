package models

import "gorm.io/gorm"

type GroupSubscription struct {
	gorm.Model
	UserId  uint   `json:"-"`
	User    User   `json:"-"`
	GroupId string `json:"-"`
	Group   Group  `json:"group"`
}
