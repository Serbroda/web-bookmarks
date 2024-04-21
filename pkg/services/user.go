package services

import (
	"context"

	"github.com/Serbroda/ragbag/pkg/dtos"
	"github.com/Serbroda/ragbag/pkg/sqlc"
)

type UserService interface {
	FindOne(ctx context.Context, id int64) (dtos.User, error)
	FindByUsername(ctx context.Context, username string) (dtos.User, error)
	Create(ctx context.Context, user CreateUser) (dtos.User, error)
}

type UserServiceSqlc struct {
	Queries *sqlc.Queries
}

type CreateUser struct {
	sqlc.CreateUserParams
	Roles []sqlc.Role
}

func (s *UserServiceSqlc) FindOne(ctx context.Context, id int64) (dtos.User, error) {
	user, err := s.Queries.FindUser(ctx, id)
	if err != nil {
		return dtos.User{}, err
	}
	u := dtos.MapUser(user)
	s.appendRoles(ctx, &u)
	return u, nil
}

func (s *UserServiceSqlc) FindByUsername(ctx context.Context, username string) (dtos.User, error) {
	user, err := s.Queries.FindUserByUsername(ctx, username)
	if err != nil {
		return dtos.User{}, err
	}
	u := dtos.MapUser(user)
	s.appendRoles(ctx, &u)
	return u, nil
}

func (s *UserServiceSqlc) Create(ctx context.Context, user CreateUser) (dtos.User, error) {
	id, err := s.Queries.CreateUser(ctx, user.CreateUserParams)
	if err != nil {
		return dtos.User{}, err
	}
	if len(user.Roles) == 0 {
		user.Roles, err = s.Queries.FindRolesByNames(ctx, []string{"USER"})
		if err != nil {
			return dtos.User{}, err
		}
	}
	for _, r := range user.Roles {
		s.Queries.InsertUserRole(ctx, sqlc.InsertUserRoleParams{
			UserID: id,
			RoleID: r.ID,
		})
	}
	return s.FindOne(ctx, id)
}

func (s *UserServiceSqlc) appendRoles(ctx context.Context, user *dtos.User) {
	roles, err := s.Queries.FindRolesByUser(ctx, user.ID)
	if err == nil {
		user.Roles = dtos.MapRoles(roles)
	}
}
