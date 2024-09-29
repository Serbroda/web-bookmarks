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
	userRepo *sqlc.Queries
}

func NewUserService(userRepo *sqlc.Queries) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(user sqlc.CreateUserParams) (int64, error) {
	if s.existsByEmail(user.Email) {
		return 0, ErrEmailAlreadyExists
	}

	if user.Username != "" {
		if s.existsUserByUsername(user.Username) {
			return 0, ErrUsernameAlreadyExists
		}
	}

	return s.userRepo.CreateUser(context.TODO(), user)
}

func (s *UserServiceImpl) GetById(id int64) (sqlc.User, error) {
	return s.userRepo.FindUserById(context.TODO(), id)
}

func (s *UserServiceImpl) GetByEmailOrUsername(emailOrUsername string) (sqlc.User, error) {
	entity, err := s.userRepo.FindUserByEmail(context.TODO(), emailOrUsername)
	if err == nil {
		return entity, nil
	}

	entity, err = s.userRepo.FindUserByUsername(context.TODO(), emailOrUsername)
	if err != nil {
		return sqlc.User{}, err
	}
	return entity, nil
}

func (s *UserServiceImpl) existsUserByUsername(username string) bool {
	if _, err := s.userRepo.FindUserByUsername(context.TODO(), username); err != nil {
		return false
	}
	return true
}

func (s *UserServiceImpl) existsByEmail(email string) bool {
	if _, err := s.userRepo.FindUserByEmail(context.TODO(), email); err != nil {
		return false
	}
	return true
}
