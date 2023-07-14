package main

import (
	"fmt"

	"github.com/Serbroda/ragbag/ui"
	"github.com/labstack/echo/v4"
)

func main() {
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
