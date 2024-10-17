package main

import (
	"backend/internal/db"
	"backend/internal/db/migrations"
	"backend/internal/db/sqlc"
	"backend/internal/http"
	"backend/internal/security"
	"backend/internal/services"
	"fmt"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	dialect = "sqlite3"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	database := db.OpenConnection(dialect, "ragbag.db")
	migrations.Migrate(database, dialect, migrations.Migrations, "sqlite")
	defer database.Close()

	queries := sqlc.New(database)

	db.InitializeData(queries)

	userService := services.NewUserService(queries)
	spaceService := services.NewSpaceService(queries)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//e.Use(middleware.Logger())

	http.RegisterAuthHandlers(e, http.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	http.RegisterUsersHandlers(api, http.UsersHandler{
		UserService: userService,
	}, "/v1")
	http.RegisterSpaceHandlers(api, http.SpaceHandler{SpaceService: spaceService}, "/v1")

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}
