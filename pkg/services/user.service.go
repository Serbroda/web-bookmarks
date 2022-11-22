package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/teris-io/shortid"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type UserService struct {
	Queries *gen.Queries
}

func NewUserService(q *gen.Queries) *UserService {
	return &UserService{Queries: q}
}

func (s *UserService) ExistsUser(ctx context.Context, username string) bool {
	exists, err := s.Queries.CountUserByName(ctx, username)
	if err != nil {
		return false
	}
	return exists > 0
}

func (s *UserService) CreateUser(ctx context.Context, arg gen.CreateUserParams) (gen.User, error) {
	return s.CreateUserWithRoles(ctx, arg, []string{})
}

func (s *UserService) CreateUserWithRoles(ctx context.Context, arg gen.CreateUserParams, roles []string) (gen.User, error) {
	e, err := s.Queries.CountUserByName(ctx, arg.Username)
	if err != nil || e > 0 {
		return gen.User{}, ErrUserAlreadyExists
	}

	shortId := shortid.MustGenerate()
	pwd, _ := utils.HashPassword(shortId)

	arg.Password = pwd
	id, err := s.Queries.CreateUser(ctx, arg)
	if err != nil {
		return gen.User{}, err
	}

	if len(roles) < 1 {
		roles = []string{"User"}
	}

	rls, err := s.Queries.FindRolesByNamesIn(ctx, strings.Join(roles, ","))
	if err != nil {
		err := s.Queries.DeleteUserFull(ctx, id)
		return gen.User{}, err
	}
	for _, r := range rls {
		s.Queries.InsertUserRole(ctx, gen.InsertUserRoleParams{
			UserID: id,
			RoleID: r.ID,
		})
	}

	return s.Queries.FindUser(ctx, id)
}
