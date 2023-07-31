package main

import (
	"fmt"

	"github.com/Serbroda/ragbag/pkg/db"
	"github.com/Serbroda/ragbag/ui"
	"github.com/labstack/echo/v4"
)

func main() {
	con := db.OpenConnection("sqlite3", "ragbag.db")
db.Migrate(con.DB, "sqlite3", db.Migrations, "migrations/sqlite")
	//db := db.OpenAndConfigure("sqlite", "ragbag.db", db.Migrations, "migrations/sqlite")

	e := echo.New()
	ui.RegisterUi(e)
	e.Logger.Fatal(e.Start(":8080"))

}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}
