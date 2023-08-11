package main

import (
	"fmt"

	"github.com/Serbroda/ragbag/cmd/server/handlers"
	"github.com/Serbroda/ragbag/pkg/db"
	"github.com/Serbroda/ragbag/pkg/security"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/sqlc"
	"github.com/Serbroda/ragbag/ui"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	dialect     = "sqlite3"
	userService services.UserService
)

func init() {
	con := db.OpenConnection(dialect, "ragbag.db")
	db.Migrate(con.DB, dialect, db.Migrations, "migrations/sqlite")
	queries := sqlc.New(con)

	userService = &services.UserServiceSqlc{Queries: queries}

	initializer := db.DataInitializer{
		Queries:     queries,
		UserService: userService,
	}
	initializer.Initialize()
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	ui.RegisterUi(e)

	baseurl := "/api"
	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{UserService: userService}, baseurl)

	baseUrlV1 := baseurl + "/v1"
	jwtMiddleware := echojwt.WithConfig(security.CreateJwtConfig(userService))
	handlers.RegisterUsersHandlers(e, handlers.UsersHandler{UserService: userService}, baseUrlV1, jwtMiddleware)

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}
