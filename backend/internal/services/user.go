package services

import (
	"backend/internal/common/random"
	"backend/internal/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserServiceImpl struct {
	userRepo *sqlc.Queries
}

func NewUserService(userRepo *sqlc.Queries) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(user sqlc.CreateUserParams) (sqlc.User, error) {
	if _, err := s.GetByEmail(user.Email); err == nil {
		return sqlc.User{}, ErrEmailAlreadyExists
	}

	if user.Username.Valid {
		i := 0
		for {
			i++

			user.Tag = sql.NullString{
				String: random.RandomStringWithCharset(4, random.CharsetAlphaUpper+random.CharsetNumbers),
				Valid:  true,
			}
			count, err := s.userRepo.CountByUsernameAndTag(context.TODO(), sqlc.CountByUsernameAndTagParams{
				LOWER: user.Username.String,
				Tag:   user.Tag,
			})
			if count == 0 {
				break
			} else if err != nil {
				fmt.Printf("username %s already exists with tag %s", user.Username, user.Tag)
			}

			log.Printf("username %s already exists with tag %s", user.Username, user.Tag)
			if i > 9999 {
				return sqlc.User{}, ErrUsernameAlreadyExists
			}
		}
	}

	id, err := s.userRepo.CreateUser(context.Background(), user)
	if err != nil {
		return sqlc.User{}, err
	}
	return s.GetById(id)
}

func (s *UserServiceImpl) GetById(id int64) (sqlc.User, error) {
	return s.userRepo.FindUserById(context.TODO(), id)
}

func (s *UserServiceImpl) GetByEmail(email string) (sqlc.User, error) {
	entity, err := s.userRepo.FindUserByEmail(context.TODO(), email)
	if err != nil {
		return sqlc.User{}, err
	}
	return entity, nil
}

func (s *UserServiceImpl) GetByUsername(username string, tag string) (sqlc.User, error) {
	entity, err := s.userRepo.FindUserByUsername(context.TODO(), sqlc.FindUserByUsernameParams{
		Username: sql.NullString{
			String: username,
			Valid:  true,
		},
		Tag: sql.NullString{
			String: tag,
			Valid:  true,
		},
	})
	if err != nil {
		return sqlc.User{}, err
	}
	return entity, nil
}
