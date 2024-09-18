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
	ErrUserAlreadyExists = errors.New("User already exists")
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
	exists, err := s.ExistsUserByUsername(user.Username)
	if err != nil {
		return err
	} else if exists {
		return ErrUserAlreadyExists
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

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.FindByUsername(context.TODO(), username)
}

func (s *UserService) ExistsUserByUsername(username string) (bool, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		} else {
			return false, err
		}
	}
	return user != nil, nil
}
