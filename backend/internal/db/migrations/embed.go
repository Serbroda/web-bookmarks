package migrations

import "embed"

var (
	//go:embed sqlite/*.sql
	Migrations embed.FS
)
