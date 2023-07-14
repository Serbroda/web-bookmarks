package ui

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed v1/dist/ragbag-ui
	FrontendDist embed.FS
	//go:embed v1/dist/ragbag-ui/index.html
	IndexHTML embed.FS
)

var (
	distDirFS     = echo.MustSubFS(FrontendDist, "v1/dist/ragbag-ui")
	distIndexHtml = echo.MustSubFS(IndexHTML, "v1/dist/ragbag-ui")
)

func RegisterUi(e *echo.Echo) {
	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)
}
