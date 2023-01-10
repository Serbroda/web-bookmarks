package sqlx

import (
	. "github.com/Serbroda/ragbag/app/pkg/models"
	"github.com/Serbroda/ragbag/app/pkg/utils"
	"github.com/jmoiron/sqlx"
	"regexp"
)

type UserServiceSqlx struct {
	DB *sqlx.DB
}

const sqlInsertUser = `
INSERT INTO users (created_at, first_name, last_name, username, password, email, active)
VALUES (CURRENT_TIMESTAMP, :first_name, :last_name, :username, :password, :email, :active)
`

const sqlFindUserById = `
SELECT *
FROM users
WHERE id = ?
LIMIT 1
`

const sqlFindUserByUsername = `
SELECT *
FROM users
WHERE lower(username) = lower(?)
LIMIT 1
`

func (s *UserServiceSqlx) CreateUser(user User) (User, error) {
	if matched, _ := regexp.MatchString(`^\$2a\$14.*$`, user.Password); !matched {
		pwd, _ := utils.HashBcrypt(user.Password)
		user.Password = pwd
	}

	result, err := s.DB.NamedExec(sqlInsertUser, &user)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return s.FindOne(id)
}

func (s *UserServiceSqlx) FindOne(id int64) (User, error) {
	var entity User
	err := s.DB.Get(&entity, sqlFindUserById, id)
	if err != nil {
		return User{}, err
	}
	return entity, nil
}

func (s *UserServiceSqlx) FindOneByUsername(username string) (User, error) {
	var entity User
	err := s.DB.Get(&entity, sqlFindUserByUsername, username)
	if err != nil {
		return User{}, err
	}
	return entity, nil
}
