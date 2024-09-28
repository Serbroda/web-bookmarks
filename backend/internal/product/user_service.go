package product

import (
	"backend/internal"
	"backend/internal/mongodb"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserServiceImpl struct {
	userRepo *mongodb.UserRepository
}

func NewUserService(userRepo *mongodb.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Create(user *internal.User) error {
	exists, err := s.existsByEmail(user.Email)
	if err != nil {
		return err
	} else if exists {
		return ErrUsernameAlreadyExists
	}

	if user.Username != "" {
		exists, err = s.existsUserByUsername(user.Username)
		if err != nil {
			return err
		} else if exists {
			return ErrUsernameAlreadyExists
		}
	}

	return s.userRepo.Save(context.TODO(), user)
}

func (s *UserServiceImpl) GetById(id string) (*internal.User, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindByID(context.TODO(), objectID)
}

func (s *UserServiceImpl) GetByEmailOrUsername(emailOrUsername string) (*internal.User, error) {
	entity, err := s.userRepo.FindByEmail(context.TODO(), emailOrUsername)
	if err == nil {
		return entity, nil
	}

	entity, err = s.userRepo.FindByUsername(context.TODO(), emailOrUsername)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *UserServiceImpl) existsUserByUsername(username string) (bool, error) {
	return s.exists(s.userRepo.FindByUsername(context.TODO(), username))
}

func (s *UserServiceImpl) existsByEmail(email string) (bool, error) {
	return s.exists(s.userRepo.FindByUsername(context.TODO(), email))
}

func (s *UserServiceImpl) exists(user *internal.User, err error) (bool, error) {
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		} else {
			return false, err
		}
	}
	return user != nil, nil
}
