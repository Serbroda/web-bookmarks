package main

import (
	"github.com/Serbroda/bookmark-manager/internal/db"
	"github.com/Serbroda/bookmark-manager/internal/server"
	"os"
)

func main() {
	run()
}

func run() {
	_, database := db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := server.NewServer(database)
	e.Logger.Fatal(e.Start(":" + port))
}
