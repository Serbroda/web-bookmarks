package services

import "github.com/Serbroda/ragbag/gen"

var (
	Service *Services
)

type Services struct {
	Queries *gen.Queries
}

func New(q *gen.Queries) *Services {
	return &Services{
		Queries: q,
	}
}
