package main

import (
	"github.com/Serbroda/bookmark-manager/internal/server"
	"os"
)

func main() {
	run()
}

func run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := server.NewServer()
	e.Logger.Fatal(e.Start(":" + port))
}
