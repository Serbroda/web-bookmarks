package services

import "github.com/Serbroda/ragbag/gen"

var (
	Services *Service
)

type Service struct {
	UserService *UserService
	RoleService *RoleService
}

func New(q *gen.Queries) *Service {
	return &Service{
		UserService: NewUserService(q),
		RoleService: NewRoleService(q),
	}
}
