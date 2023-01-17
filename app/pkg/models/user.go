package models

import (
	"time"
)

type User struct {
	BaseModel
	FirstName             string     `db:"first_name" json:"first_name"`
	LastName              string     `db:"last_name" json:"last_name"`
	Username              string     `db:"username" json:"username"`
	Password              string     `db:"password" json:"-"`
	Email                 string     `db:"email" json:"email"`
	Active                bool       `db:"active" json:"-"`
	ActivationConfirmedAt *time.Time `db:"activation_confirmed_at" json:"-"`
	Roles                 []Role     `db:"-" json:"roles"`
}
