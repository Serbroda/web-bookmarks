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
	ErrActivationCodeExpired = errors.New("activation code expired")
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

	if matched, _ := regexp.MatchString(`^\$2a\$14.*$`, arg.Password); !matched {
		pwd, _ := utils.HashBcrypt(arg.Password)
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

func (s *Services) FindActivationToken(ctx context.Context, token string) (gen.ActivationToken, error) {
	at, err := s.Queries.FindActivationCode(ctx, utils.HashSha3256(token))
	if err != nil {
		return gen.ActivationToken{}, ErrUserNotFound
	}
	return at, nil
}

func (s *Services) ChangePassword(ctx context.Context, userId int64, password string) error {
	user, err := s.FindUser(ctx, userId)
	if err != nil {
		return err
	}

	if matched, _ := regexp.MatchString(`^\$2a\$14.*$`, password); !matched {
		pwd, _ := utils.HashBcrypt(password)
		password = pwd
	}

	err = s.Queries.UpdateUser(ctx, gen.UpdateUserParams{
		ID:       user.ID,
		Password: password,

		ActivationConfirmedAt: user.ActivationConfirmedAt,
		Active:                user.Active,
		FirstName:             user.FirstName,
		LastName:              user.LastName,
		Email:                 user.Email,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Services) ActivateUser(ctx context.Context, token string) error {
	at, err := s.FindActivationToken(ctx, token)
	if err != nil {
		return err
	}

	user, err := s.FindUser(ctx, at.UserID)
	if err != nil {
		return err
	}

	if user.Active {
		return ErrUserAlreadyActive
	}

	if at.ExpiresAt.Valid && at.ExpiresAt.Time.Before(time.Now()) {
		return ErrActivationCodeExpired
	}

	err = s.Queries.UpdateUser(ctx, gen.UpdateUserParams{
		ID:                    user.ID,
		ActivationConfirmedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Active:                true,

		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
	})

	if err != nil {
		return err
	}

	_, err = s.CreateSpace(ctx, gen.CreateSpaceParams{
		ShortID:    shortid.MustGenerate(),
		OwnerID:    user.ID,
		Name:       fmt.Sprintf("%s's Space", user.FirstName),
		Visibility: "PRIVATE",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Services) CreateActivationToken(ctx context.Context, userId int64) (string, error) {
	activationToken := utils.RandomString(64)

	err := s.Queries.InsertActivationToken(ctx, gen.InsertActivationTokenParams{
		UserID:    userId,
		TokenHash: utils.HashSha3256(activationToken),
		ExpiresAt: sql.NullTime{Time: time.Now().Add(time.Hour * 48), Valid: true},
	})

	if err != nil {
		return "", err
	}

	return activationToken, nil
}

func (s *Services) CreatePasswordResetToken(ctx context.Context, userId int64) (string, error) {
	activationToken := utils.RandomString(64)

	err := s.Queries.InsertPasswordResetToken(ctx, gen.InsertPasswordResetTokenParams{
		UserID:    userId,
		TokenHash: utils.HashSha3256(activationToken),
		ExpiresAt: time.Now().Add(time.Hour * 4),
	})

	if err != nil {
		return "", err
	}

	return activationToken, nil
}
