package main

import (
	"backend/internal/db"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type Person struct {
	Name string
	Age  int
	City string
}

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	checkDB()

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func checkDB() {
	database, client, ctx, cancel := db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection(client, ctx, cancel)

	collection := database.Collection("people")

	// Insert a document
	res, err := collection.InsertOne(context.Background(), Person{
		"John Doe",
		30,
		"New York",
	})

	fmt.Printf("Inserted a single document: %v\n", res.InsertedID)

	// Find a document
	var result Person
	err = collection.FindOne(context.Background(), bson.D{{"name", "John Doe"}}).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Println("No Document found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Found document: %+v\n", result)
	}
}
