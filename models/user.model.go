package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
