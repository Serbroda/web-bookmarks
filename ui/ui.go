package ui

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed v1/dist
	FrontendDist embed.FS
	//go:embed v1/dist/index.html
	IndexHTML embed.FS
)

var (
	distDirFS     = echo.MustSubFS(FrontendDist, "v1/dist")
	distIndexHtml = echo.MustSubFS(IndexHTML, "v1/dist")
)

func RegisterUi(e *echo.Echo) {
	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)
}
