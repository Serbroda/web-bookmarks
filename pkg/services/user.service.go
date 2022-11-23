package services

import (
	"context"
	"errors"
	"regexp"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/teris-io/shortid"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func (s *Services) ExistsUser(ctx context.Context, username string) bool {
	exists, err := s.Queries.CountUserByName(ctx, username)
	if err != nil {
		return false
	}
	return exists > 0
}

func (s *Services) CreateUser(ctx context.Context, arg gen.CreateUserParams) (gen.User, error) {
	return s.CreateUserWithRoles(ctx, arg, []string{})
}

func (s *Services) CreateUserWithRoles(ctx context.Context, arg gen.CreateUserParams, roles []string) (gen.User, error) {
	if s.ExistsUser(ctx, arg.Username) {
		return gen.User{}, ErrUserAlreadyExists
	}

	shortId := shortid.MustGenerate()

	if matched, _ := regexp.MatchString(`^\$2a\$14.*$`, arg.Password); !matched {
		pwd, _ := utils.HashPassword(shortId)
		arg.Password = pwd
	}

	id, err := s.Queries.CreateUser(ctx, arg)
	if err != nil {
		return gen.User{}, err
	}

	if len(roles) < 1 {
		roles = []string{"USER"}
	}

	s.FindRolesByNamesIn(ctx, roles)

	for _, r := range s.FindRolesByNamesIn(ctx, roles) {
		s.Queries.InsertUserRole(ctx, gen.InsertUserRoleParams{
			UserID: id,
			RoleID: r.ID,
		})
	}

	return s.Queries.FindUser(ctx, id)
}
