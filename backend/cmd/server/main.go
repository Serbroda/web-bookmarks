package main

import (
	"backend/internal/db"
	"backend/internal/db/migrations"
	"backend/internal/http"
	"backend/internal/security"
	"backend/internal/services"
	"backend/internal/sqlc"
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
	db := db.OpenConnection(dialect, "ragbag.db")
	migrations.Migrate(db, dialect, migrations.Migrations, "sqlite")
	defer db.Close()

	queries := sqlc.New(db)

	userService := services.NewUserService(queries)
	//contentService := services.NewContentService(spaceRepo, pageRepo, bookmarkRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	http.RegisterAuthHandlers(e, http.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	http.RegisterUsersHandlers(api, http.UsersHandler{}, "/v1")
	//http.RegisterSpaceHandlers(api, http.SpaceHandler{ContentService: contentService}, "/v1")

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}
