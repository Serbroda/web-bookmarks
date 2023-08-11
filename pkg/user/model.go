package user

import (
	"strings"

	"github.com/Serbroda/ragbag/pkg/sqlc"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Roles     []Role `json:"roles,omitempty"`
}

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (u *User) IsAdmin() bool {
	return u.HasAnyRole("admin")
}

func (u *User) HasAnyRole(roles ...string) bool {
	for _, ur := range u.Roles {
		for _, r := range roles {
			if strings.EqualFold(ur.Name, r) {
				return true
			}
		}
	}
	return false
}

func (u *User) RolesAsStrings() []string {
	roles := []string{}
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}
	return roles
}

func MapUser(entity sqlc.User) User {
	return User{
		ID:        entity.ID,
		Username:  entity.Username,
		Password:  entity.Password,
		Email:     entity.Email,
		FirstName: entity.FirstName.String,
		LastName:  entity.LastName.String,
	}
}

func MapRole(entity sqlc.Role) Role {
	return Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func MapRoles(entities []sqlc.Role) []Role {
	var roles []Role
	for _, r := range entities {
		roles = append(roles, MapRole(r))
	}
	return roles
}
