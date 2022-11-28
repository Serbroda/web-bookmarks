package services

import "github.com/Serbroda/ragbag/gen"

var (
	Service *Services
)

type Services struct {
	Queries *gen.Queries
}

func New(q *gen.Queries) *Services {
	Service = &Services{
		Queries: q,
	}
	return Service
}
