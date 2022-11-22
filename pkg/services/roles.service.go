package services

import (
	"context"

	"github.com/Serbroda/ragbag/gen"
)

type RoleService struct {
	Queries *gen.Queries
}

func NewRoleService(q *gen.Queries) *RoleService {
	return &RoleService{Queries: q}
}

func (s *RoleService) FindRolesByNamesIn(ctx context.Context, roles []string) ([]gen.Role, error) {
	var result []gen.Role
	for _, r := range roles {
		res, err := s.Queries.FindRoleByName(ctx, r)
		if err == nil && res.ID > 0 {
			result = append(result, res)
		}
	}
	return result, nil
}
