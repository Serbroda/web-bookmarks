package main

import (
	"backend/cmd/server/handlers"
	"backend/internal/db"
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	_, database := db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	tryDB(database)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{}, "")

	//baseUrlV1 := "/api/v1"
	//jwtMiddleware := echojwt.WithConfig(security.CreateJwtConfig(userService))
	//handlers.RegisterUsersHandlers(e, handlers.UsersHandler{UserService: userService}, baseUrlV1, jwtMiddleware)

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func tryDB(database *mongo.Database) {
	spaceRepo := repository.NewSpaceRepository(database.Collection("spaces"))

	// Neues Space-Objekt erstellen
	newSpace := &model.Space{
		Name:    "New Space",
		OwnerID: bson.NewObjectID(),
	}

	// Speichern (Insert oder Update)
	err := spaceRepo.Save(context.TODO(), newSpace)
	if err != nil {
		log.Fatalf("Failed to save space: %v", err)
	}

	fmt.Printf("Space saved with ID: %s\n", newSpace.ID.Hex())

	found, err := spaceRepo.FindByID(context.TODO(), newSpace.ID)
	if err != nil {
		log.Fatalf("Failed to save space: %v", err)
	}

	fmt.Printf("Found : %v\n", found)

	found.Name = "New name"
	err = spaceRepo.Save(context.TODO(), found)
	if err != nil {
		log.Fatalf("Failed to save space: %v", err)
	}

	found, err = spaceRepo.FindByID(context.TODO(), newSpace.ID)
	if err != nil {
		log.Fatalf("Failed to save space: %v", err)
	}

	fmt.Printf("Found : %v\n", found)
}
