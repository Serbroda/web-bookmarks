package services

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetByEmailOrUsername(ctx context.Context, email string) (*models.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (u userService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetByEmailOrUsername(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
