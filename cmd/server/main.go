package main

import (
	"context"
	"fmt"

	"github.com/Serbroda/ragbag/cmd/server/handlers"
	"github.com/Serbroda/ragbag/pkg/db"
	"github.com/Serbroda/ragbag/pkg/sqlc"
	"github.com/Serbroda/ragbag/pkg/user"
	"github.com/Serbroda/ragbag/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dialect = "sqlite3"

func main() {
	con := db.OpenConnection(dialect, "ragbag.db")
	db.Migrate(con.DB, dialect, db.Migrations, "migrations/sqlite")
	queries := sqlc.New(con)

	us := user.UserServiceSqlc{Queries: queries}
	fmt.Println(us.FindOne(context.Background(), 1))

	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	ui.RegisterUi(e)

	baseurl := "/api"
	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{UserService: &us}, baseurl)

	/*jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &controllers.JwtCustomClaims{},
		SigningKey: []byte(jwtSecretKey),
	})*/

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}
