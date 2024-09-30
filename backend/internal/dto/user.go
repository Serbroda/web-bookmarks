package dto

import (
	"backend/internal/sqlc"
	"fmt"
)

type UserDto struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Tag      string `db:"tag" json:"tag"`
}

func (d *UserDto) UsernameWithTag() string {
	return fmt.Sprintf("%s:%s", d.Username, d.Tag)
}

func UserDtoFromUser(user sqlc.User) UserDto {
	return UserDto{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username.String,
		Tag:      user.Tag.String,
	}
}
