package sqlx

import (
	"database/sql"
	"errors"
	"github.com/Serbroda/ragbag/pkg/model"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
)

type UserServiceSqlx struct {
	DB *sqlx.DB
}

func (u *UserServiceSqlx) embedRoles(user *model.User) {
	if user == nil {
		return
	}
	var roles []model.Role
	err := u.DB.Select(&roles, "SELECT r.* FROM roles r INNER JOIN users_roles ur on ur.role_id = r.id WHERE ur.user_id = $1", user.Id)
	if err != nil {
		return
	}
	user.Roles = roles
}

func (u *UserServiceSqlx) FindOne(id int64) (*model.User, error) {
	var entity model.User
	err := u.DB.Get(&entity, "SELECT * FROM users WHERE id = $1 deleted_at is null", id)
	if err != nil {
		return nil, err
	}
	u.embedRoles(&entity)
	return &entity, nil
}

func (u *UserServiceSqlx) FindOneByUsername(username string) (*model.User, error) {
	var entity model.User
	err := u.DB.Get(&entity, "SELECT * FROM users WHERE lower(username) = lower($1) deleted_at is null", username)
	if err != nil {
		return nil, err
	}
	u.embedRoles(&entity)
	return &entity, nil
}

func (u *UserServiceSqlx) FindOneByEmail(email string) (*model.User, error) {
	var entity model.User
	err := u.DB.Get(&entity, "SELECT * FROM users WHERE lower(email) = lower($1) deleted_at is null", email)
	if err != nil {
		return nil, err
	}
	u.embedRoles(&entity)
	return &entity, nil
}

func (u *UserServiceSqlx) Create(params model.UserCRUD) (*model.User, error) {
	if user, err := u.FindOneByUsername(params.Username); exists(user, err) {
		return nil, ErrUsernameAlreadyExists
	}
	if user, err := u.FindOneByEmail(params.Email); exists(user, err) {
		return nil, ErrEmailAlreadyExists
	}
	res, err := u.DB.Exec(`INSERT INTO users (created_at, first_name, last_name, username, password, email, active)
		VALUES (CURRENT_TIMESTAMP, $1, $2, $3, $4, $5, $6)`, params.FirstName, params.LastName, params.Username, params.Password, params.Email, params.Active)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return u.FindOne(id)
}

func exists(user *model.User, err error) bool {
	return (err != nil && err == sql.ErrNoRows) || user == nil
}
