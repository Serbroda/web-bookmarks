package services

import (
	"backend/internal/sqlc"
	"context"
	"errors"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserServiceImpl struct {
	queries *sqlc.Queries
}

func NewUserService(userRepo *sqlc.Queries) *UserServiceImpl {
	return &UserServiceImpl{
		queries: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(user sqlc.CreateUserParams) (sqlc.User, error) {
	count, err := s.queries.CountUserByEmail(context.TODO(), user.Username)
	if err != nil {
		return sqlc.User{}, err
	} else if count > 0 {
		return sqlc.User{}, ErrEmailAlreadyExists
	}

	count, err = s.queries.CountUserByUsername(context.TODO(), user.Username)
	if err != nil {
		return sqlc.User{}, err
	} else if count > 0 {
		return sqlc.User{}, ErrUsernameAlreadyExists
	}

	id, err := s.queries.CreateUser(context.Background(), user)
	if err != nil {
		return sqlc.User{}, err
	}
	return s.GetById(id)
}

func (s *UserServiceImpl) GetById(id int64) (sqlc.User, error) {
	return s.queries.FindUserById(context.TODO(), id)
}

func (s *UserServiceImpl) GetByEmailOrUsername(emailOrUsername string) (sqlc.User, error) {
	entity, err := s.queries.FindUserByEmailOrUsername(context.TODO(), sqlc.FindUserByEmailOrUsernameParams{
		Email:    emailOrUsername,
		Username: emailOrUsername,
	})
	if err != nil {
		return sqlc.User{}, err
	}
	return entity, nil
}
