package services

import (
	"errors"
	. "github.com/Serbroda/ragbag/app/pkg/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService interface {
	CreateUser(user User) (User, error)
	FindOne(id int64) (User, error)
	FindOneByUsername(username string) (User, error)
}
