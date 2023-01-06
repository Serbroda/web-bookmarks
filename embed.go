package ragbag

import (
	"embed"
)

var (
	//go:embed app/resources/db/migrations/sqlite/*.sql
	Migrations embed.FS
	//go:embed frontend/dist
	FrontendDist embed.FS
	//go:embed frontend/dist/index.html
	IndexHTML embed.FS
)
