package user

import "github.com/Serbroda/ragbag/pkg/sqlc"

type User struct {
	ID    int64 `json:"id"`
	Roles []Role
}

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func MapUser(entity sqlc.User) User {
	return User{
		ID: entity.ID,
	}
}

func MapRole(entity sqlc.Role) Role {
	return Role{
		ID: entity.ID,
	}
}

func MapRoles(entities []sqlc.Role) []Role {
	var roles []Role
	for _, r := range entities {
		roles = append(roles, MapRole(r))
	}
	return roles
}
