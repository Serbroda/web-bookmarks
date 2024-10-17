package services

import (
	sqlc2 "backend/internal/db/sqlc"
	"context"
	"errors"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserServiceImpl struct {
	queries *sqlc2.Queries
}

func NewUserService(userRepo *sqlc2.Queries) *UserServiceImpl {
	return &UserServiceImpl{
		queries: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(user sqlc2.CreateUserParams) (sqlc2.User, error) {
	count, err := s.queries.CountUserByEmail(context.TODO(), user.Username)
	if err != nil {
		return sqlc2.User{}, err
	} else if count > 0 {
		return sqlc2.User{}, ErrEmailAlreadyExists
	}

	count, err = s.queries.CountUserByUsername(context.TODO(), user.Username)
	if err != nil {
		return sqlc2.User{}, err
	} else if count > 0 {
		return sqlc2.User{}, ErrUsernameAlreadyExists
	}

	entity, err := s.queries.CreateUser(context.Background(), user)
	if err != nil {
		return sqlc2.User{}, err
	}
	return entity, nil
}

func (s *UserServiceImpl) GetUserById(id int64) (sqlc2.User, error) {
	return s.queries.FindUserById(context.TODO(), id)
}

func (s *UserServiceImpl) GetUserByEmailOrUsername(emailOrUsername string) (sqlc2.User, error) {
	entity, err := s.queries.FindUserByEmailOrUsername(context.TODO(), sqlc2.FindUserByEmailOrUsernameParams{
		Email:    emailOrUsername,
		Username: emailOrUsername,
	})
	if err != nil {
		return sqlc2.User{}, err
	}
	return entity, nil
}
