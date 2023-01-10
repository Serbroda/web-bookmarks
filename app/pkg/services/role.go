package services

import (
	. "github.com/Serbroda/ragbag/app/pkg/models"
)

type RoleService interface {
	FindRoleByName(name string) (Role, error)
	FindRolesByNames(names ...string) []Role
	FindUserRoles(userId int64) []Role
	InsertUserRole(userId, roleId int64) error
}
