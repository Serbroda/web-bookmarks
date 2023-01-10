package sqlx

import (
	. "github.com/Serbroda/ragbag/app/pkg/models"
	"github.com/jmoiron/sqlx"
)

type RoleServiceSqlx struct {
	DB *sqlx.DB
}

const sqlFindRoleByName = `
SELECT *
FROM roles
WHERE lower(name) = lower(?)
LIMIT 1
`

const sqlFindUserRoles = `
SELECT r.*
FROM roles r
	INNER JOIN users_roles ur on r.id = ur.role_id
WHERE ur.user_id = ?
`

const sqlInsertUserRole = `
INSERT INTO users_roles(user_id, role_id, created_at) 
VALUES (?, ?, CURRENT_TIMESTAMP)
`

func (s *RoleServiceSqlx) FindRoleByName(name string) (Role, error) {
	var entity Role
	err := s.DB.Get(&entity, sqlFindRoleByName, name)
	if err != nil {
		return Role{}, err
	}
	return entity, nil
}

func (s *RoleServiceSqlx) FindRolesByNames(names ...string) []Role {
	var result []Role
	for _, r := range names {
		res, err := s.FindRoleByName(r)
		if err == nil && res.ID > 0 {
			result = append(result, res)
		}
	}
	return result
}

func (s *RoleServiceSqlx) FindUserRoles(userId int64) []Role {
	var result []Role
	s.DB.Select(&result, sqlFindUserRoles, userId)
	return result
}

func (s *RoleServiceSqlx) InsertUserRole(userId, roleId int64) error {
	_, err := s.DB.Exec(sqlInsertUserRole, userId, roleId)
	return err
}
