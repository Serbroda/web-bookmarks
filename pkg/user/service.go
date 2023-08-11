package user

import (
	"context"

	"github.com/Serbroda/ragbag/pkg/sqlc"
)

type UserService interface {
	FindOne(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user sqlc.CreateUserParams) (User, error)
}

type UserServiceSqlc struct {
	Queries *sqlc.Queries
}

func (s *UserServiceSqlc) FindOne(ctx context.Context, id int64) (User, error) {
	user, err := s.Queries.FindUser(ctx, id)
	if err != nil {
		return User{}, err
	}
	u := MapUser(user)
	s.appendRoles(ctx, &u)
	return u, nil
}

func (s *UserServiceSqlc) FindByUsername(ctx context.Context, username string) (User, error) {
	user, err := s.Queries.FindUserByUsername(ctx, username)
	if err != nil {
		return User{}, err
	}
	u := MapUser(user)
	s.appendRoles(ctx, &u)
	return u, nil
}

func (s *UserServiceSqlc) Create(ctx context.Context, user sqlc.CreateUserParams) (User, error) {
	id, err := s.Queries.CreateUser(ctx, user)
	if err != nil {
		return User{}, err
	}
	return s.FindOne(ctx, id)
}

func (s *UserServiceSqlc) appendRoles(ctx context.Context, user *User) {
	roles, err := s.Queries.FindRolesByUser(ctx, user.ID)
	if err == nil {
		user.Roles = MapRoles(roles)
	}
}
