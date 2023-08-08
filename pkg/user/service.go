package user

import (
	"context"

	"github.com/Serbroda/ragbag/pkg/sqlc"
)

type UserService interface {
	FindOne(id int64) User
}

type UserServiceSqlc struct {
	Queries *sqlc.Queries
}

func (s *UserServiceSqlc) FindOne(ctx context.Context, id int64) User {
	user, err := s.Queries.FindUser(ctx, id)
	if err != nil {
		return User{}
	}
	u := MapUser(user)
	s.appendRoles(ctx, &u)
	return u
}

func (s *UserServiceSqlc) appendRoles(ctx context.Context, user *User) {
	roles, err := s.Queries.FindRolesByUser(ctx, user.ID)
	if err == nil {
		user.Roles = MapRoles(roles)
	}
}
