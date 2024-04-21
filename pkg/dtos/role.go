package dtos

import "github.com/Serbroda/ragbag/pkg/sqlc"

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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

func MapMany[T any]() []T {

}
