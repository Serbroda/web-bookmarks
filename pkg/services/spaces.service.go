package services

import (
	"context"

	"github.com/Serbroda/ragbag/gen"
)

func (s *Services) CreateSpace(ctx context.Context, params gen.CreateSpaceParams) (gen.Space, error) {
	id, err := s.Queries.CreateSpace(ctx, params)
	if err != nil {
		return gen.Space{}, err
	}
	space, err := s.Queries.FindSpace(ctx, id)
	if err != nil {
		return gen.Space{}, err
	}
	return space, nil
}
