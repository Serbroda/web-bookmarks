package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/teris-io/shortid"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUserAlreadyActive     = errors.New("user already active")
	ErrActivationCodeExpores = errors.New("activation code expired")
)

func (s *Services) ExistsUser(ctx context.Context, username string) bool {
	_, err := s.FindUserByUsername(ctx, username)
	return err == nil
}

func (s *Services) FindUser(ctx context.Context, id int64) (gen.User, error) {
	user, err := s.Queries.FindUser(ctx, id)
	if err != nil || user.ID < 1 {
		return gen.User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *Services) FindUserByUsername(ctx context.Context, username string) (gen.User, error) {
	user, err := s.Queries.FindUserByUsername(ctx, username)
	if err != nil || user.ID < 1 {
		return gen.User{}, ErrUserNotFound
	}
	return user, nil
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

	for _, r := range s.FindRolesByNamesIn(ctx, roles) {
		if !s.HasUserRole(ctx, id, r.Name) {
			s.Queries.InsertUserRole(ctx, gen.InsertUserRoleParams{
				UserID: id,
				RoleID: r.ID,
			})
		}
	}

	return s.FindUser(ctx, id)
}

func (s *Services) HasUserRole(ctx context.Context, id int64, role string) bool {
	res, err := s.Queries.CountUserRole(ctx, gen.CountUserRoleParams{
		UserID: id,
		Name:   role,
	})
	return err != nil && res > 0
}

func (s *Services) FindUserByActivationCode(ctx context.Context, code string) (gen.User, error) {
	user, err := s.Queries.FindUserByActivationCode(ctx, sql.NullString{String: code, Valid: true})
	if err != nil || user.ID < 1 {
		fmt.Printf("error: %v, user: %v", err, user)
		return gen.User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *Services) ActivateUser(ctx context.Context, code string) error {
	user, err := s.FindUserByActivationCode(ctx, code)
	if err != nil {
		return err
	}

	if user.Active {
		return ErrUserAlreadyActive
	}

	if !user.ActivationCodeExpiresAt.Valid || user.ActivationCodeExpiresAt.Time.Before(time.Now()) {
		return ErrActivationCodeExpores
	}

	err = s.Queries.UpdateUser(ctx, gen.UpdateUserParams{
		ID:                    user.ID,
		ActivationConfirmedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Active:                true,

		FirstName:               user.FirstName,
		LastName:                user.LastName,
		Password:                user.Password,
		Name:                    user.Name,
		Email:                   user.Email,
		ActivationCode:          user.ActivationCode,
		ActivationSentAt:        user.ActivationSentAt,
		ActivationCodeExpiresAt: user.ActivationCodeExpiresAt,
	})

	if err != nil {
		return err
	}

	_, err = s.CreateSpace(ctx, gen.CreateSpaceParams{
		ShortID:    shortid.MustGenerate(),
		OwnerID:    user.ID,
		Name:       "Default",
		Visibility: "PRIVATE",
	})

	if err != nil {
		return err
	}

	return nil
}
