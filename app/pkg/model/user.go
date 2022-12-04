package model

import "time"

type User struct {
	BaseModel
	FirstName             string     `db:"first_name" json:"firstName"`
	LastName              string     `db:"last_name" json:"lastName"`
	Username              string     `db:"username" json:"username"`
	Password              string     `db:"password" json:"password"`
	Email                 string     `db:"email" json:"email"`
	Active                bool       `db:"active" json:"-"`
	ActivationConfirmedAt *time.Time `db:"activation_confirmed_at" json:"-"`
	Roles                 []Role
}

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UserCRUD struct {
	NewUser
	Active                bool
	ActivationConfirmedAt *time.Time
}

type UserService interface {
	FindOne(id int64) (*User, error)
	FindOneByUsername(username string) (*User, error)
	FindOneByEmail(email string) (*User, error)
	Create(params UserCRUD) (*User, error)
}
