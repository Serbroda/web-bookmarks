package main

import (
	"backend/internal/db"
	"backend/internal/model"
	"backend/internal/repository"
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
	db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	checkDB2(db.Database)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func checkDB2(db *mongo.Database) {
	spaceRepo := repository.NewMongoSpaceRepository(db.Collection("spaces"))
	pageRepo := repository.NewMongoPageRepository(db.Collection("pages"))

	// Ein neues Space-Dokument einfügen
	space := model.Space{
		Name:        "Development Resources",
		Description: "A space for developers",
	}
	err := spaceRepo.Save(context.TODO(), &space)
	if err != nil {
		log.Fatal("Failed to insert space:", err)
	}

	fmt.Println("Space successfully inserted")

	// Space anhand des Namens suchen
	foundSpace, err := spaceRepo.FindBySpaceName(context.TODO(), "Development Resources")
	if err != nil {
		log.Fatal("Failed to find space:", err)
	}
	fmt.Println("Found space:", foundSpace)

	foundSpace.Name = "Danny"
	err = spaceRepo.Save(context.TODO(), foundSpace)
	if err != nil {
		log.Fatal("Failed to update space space:", err)
	}
	// Space anhand des Namens suchen
	foundSpace, err = spaceRepo.FindBySpaceName(context.TODO(), "Danny")
	if err != nil {
		log.Fatal("Failed to find space:", err)
	}
	fmt.Println("Found space:", foundSpace)

	space2 := model.Space{
		Name:        "222",
		Description: "222",
	}
	err = spaceRepo.Save(context.TODO(), &space2)
	if err != nil {
		log.Fatal("Failed to insert space:", err)
	}

	page := model.Page{
		Name:    "Development Resources2",
		SpaceID: foundSpace.ID,
	}
	err = pageRepo.Save(context.TODO(), &page)
	if err != nil {
		log.Fatal("Failed to insert page:", err)
	}

	fmt.Println("Page successfully inserted")
}

func checkDB(database *mongo.Database) {
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
