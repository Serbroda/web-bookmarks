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

	for _, r := range s.getRoles(ctx, roles) {
		s.Queries.InsertUserRole(ctx, gen.InsertUserRoleParams{
			UserID: id,
			RoleID: r.ID,
		})
	}

	return s.Queries.FindUser(ctx, id)
}

func (s *UserService) getRoles(ctx context.Context, roleNames []string) []gen.Role {
	var roles []gen.Role
	for _, r := range roleNames {
		role, err := s.Queries.FindRoleByName(ctx, r)
		if err == nil {
			roles = append(roles, role)
		}
	}
	return roles
}
