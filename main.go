package main

import (
	"embed"
	"fmt"

	"github.com/Serbroda/ragbag/gen/public"
	gen "github.com/Serbroda/ragbag/gen/restricted"
	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/handlers"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teris-io/shortid"
)

var (
	//go:embed all:frontend/dist
	dist embed.FS
	//go:embed frontend/dist/index.html
	indexHTML     embed.FS
	distDirFS     = echo.MustSubFS(dist, "frontend/dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "frontend/dist")
)

var (
	version string
)

func main() {
	fmt.Println("version=", version)

	var dbName = utils.GetEnv("DB_NAME", "ragbag.db")
	var serverAddress = utils.GetEnv("SERVER_URL", "0.0.0.0:8080")

	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	shortid.SetDefault(sid)

	database.Connect(database.ConnectionOptions{Name: dbName})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerHandlers(e)

	e.Logger.Fatal(e.Start(serverAddress))
}

func registerHandlers(e *echo.Echo) {
	registerStaticHandlers(e)
	registerApiHandlers(e)
}

func registerStaticHandlers(e *echo.Echo) {
	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)
}

func registerApiHandlers(e *echo.Echo) {
	var publicApi handlers.PublicServerInterfaceImpl
	public.RegisterHandlersWithBaseURL(e, &publicApi, "/api")

	api := e.Group("/api/v1")
	config := middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))

	var restrictedApi handlers.RestrictedServerInterfaceImpl
	gen.RegisterHandlers(api, &restrictedApi)
}
