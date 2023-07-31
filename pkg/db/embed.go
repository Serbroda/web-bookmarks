package db

import "embed"

var (
	//go:embed migrations/sqlite/*.sql
	Migrations embed.FS
)
