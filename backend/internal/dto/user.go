package dto

import (
	"backend/internal/sqlc"
)

type UserDto struct {
	ID          int64   `json:"id"`
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	DisplayName *string `json:"displayName"`
}

func UserDtoFromUser(user sqlc.User) UserDto {
	return UserDto{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}
}
