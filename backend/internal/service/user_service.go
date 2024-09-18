package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	exists, err := s.ExistsByEmail(user.Email)
	if err != nil {
		return err
	} else if exists {
		return ErrUsernameAlreadyExists
	}

	if user.Username != "" {
		exists, err = s.ExistsUserByUsername(user.Username)
		if err != nil {
			return err
		} else if exists {
			return ErrUsernameAlreadyExists
		}
	}

	return s.userRepo.Save(context.TODO(), user)
}

func (s *UserService) GetUserById(id string) (*model.User, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByID(context.TODO(), objectID)
}

func (s *UserService) GetUserByEmailOrUsername(emailOrUsername string) (*model.User, error) {
	user, err := s.GetUserByEmail(emailOrUsername)
	if err != nil {
		user, err = s.GetUserByUsername(emailOrUsername)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.FindByUsername(context.TODO(), username)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.FindByEmail(context.TODO(), email)
}

func (s *UserService) ExistsUserByUsername(username string) (bool, error) {
	return s.exists(s.GetUserByUsername(username))
}

func (s *UserService) ExistsByEmail(username string) (bool, error) {
	return s.exists(s.GetUserByEmail(username))
}

func (s *UserService) exists(user *model.User, err error) (bool, error) {
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		} else {
			return false, err
		}
	}
	return user != nil, nil
}
