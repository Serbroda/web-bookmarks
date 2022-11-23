package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/gen/public"
	"github.com/Serbroda/ragbag/gen/restricted"
	"github.com/Serbroda/ragbag/pkg/db"
	"github.com/Serbroda/ragbag/pkg/handlers"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teris-io/shortid"
)

var (
	//go:embed resources/db/migrations/*.sql
	migrations embed.FS

	//go:embed all:frontend/dist
	dist embed.FS
	//go:embed frontend/dist/index.html
	indexHTML     embed.FS
	distDirFS     = echo.MustSubFS(dist, "frontend/dist")
	distIndexHtml = echo.MustSubFS(indexHTML, "frontend/dist")
)

var (
	version       string
	serverAddress string = utils.GetEnvFallback("SERVER_URL", "0.0.0.0:8080")
	dbAddress     string = utils.MustGetEnv("DB_ADDRESS")
	dbName        string = utils.GetEnvFallback("DB_NAME", "ragbag")
	dbUser        string = utils.GetEnvFallback("DB_USER", "ragbag")
	dbPassword    string = utils.MustGetEnv("DB_PASSWORD")
	jwtSecretKey  string = utils.MustGetEnv("JWT_SECRET_KEY")
)

func main() {
	fmt.Println("version=", version)

	db.OpenAndConfigure("mysql", getDsn(dbUser, dbPassword, dbAddress, dbName), migrations, "resources/db/migrations")

	services := services.New(db.Queries)
	db.InitializeAdmin(context.Background(), services)

	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	shortid.SetDefault(sid)

	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	registerHandlers(e, db.Queries, services)

	e.Logger.Fatal(e.Start(serverAddress))
}

func registerHandlers(e *echo.Echo, queries *gen.Queries, services *services.Services) {
	registerStaticHandlers(e)
	registerApiHandlers(e, queries, services)
}

func registerStaticHandlers(e *echo.Echo) {
	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)
}

func registerApiHandlers(e *echo.Echo, queries *gen.Queries, services *services.Services) {
	baseUrl := "/api/v1"

	var publicApi = handlers.PublicServerInterfaceImpl{
		Services: services,
		Queries:  queries,
	}
	public.RegisterHandlersWithBaseURL(e, &publicApi, baseUrl)

	api := e.Group(baseUrl)
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte(jwtSecretKey),
	}))

	var restrictedApi handlers.RestrictedServerInterfaceImpl
	restricted.RegisterHandlers(api, &restrictedApi)
}

func getDsn(user, password, address, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, address, database)
}
