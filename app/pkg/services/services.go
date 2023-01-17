package services

import . "github.com/Serbroda/ragbag/app/sqlc"

var (
	Service *Services
)

type Services struct {
	Queries *Queries
}

func NewServices(q *Queries) *Services {
	Service = &Services{
		Queries: q,
	}
	return Service
}
