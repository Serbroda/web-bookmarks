package main

import (
	"backend/cmd/server/handlers"
	"backend/internal/db"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/security"
	"backend/internal/service"
	"context"
	"errors"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	_, database := db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	userRepo := repository.NewUserRepository(database.Collection("users"))
	userService := service.NewUserService(userRepo)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	handlers.RegisterUsersHandlers(api, handlers.UsersHandler{UserService: userService}, "/v1")

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
	userRepo := repository.NewUserRepository(database.Collection("users"))
	pageRepo := repository.NewPageRepository(database.Collection("pages"))
	spaceRepo := repository.NewSpaceRepository(database.Collection("spaces"))
	bookmarkRepo := repository.NewBookmarkRepository(database.Collection("bookmarks"))

	name := "admin"
	user, err := userRepo.FindByUsername(context.TODO(), name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			user = &model.User{
				Username: name,
			}
			userRepo.Save(context.Background(), user)
			fmt.Println("created user")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Found user with username: %v\n", user.Username)
	}

	space := model.Space{
		Name: "Test Space",
	}
	spaceRepo.Save(context.TODO(), &space)

	page := model.Page{
		Name:    "Test Page",
		SpaceID: space.ID,
	}
	pageRepo.Save(context.TODO(), &page)

	bookmark := model.Bookmark{
		Title:  "Test Bookmark",
		URL:    "http://google.de",
		PageId: page.ID,
	}
	bookmarkRepo.Save(context.TODO(), &bookmark)

	// Search
	spaces, err := spaceRepo.FindAll(context.TODO())
	if err == nil {
		for _, space := range spaces {
			fmt.Printf(" - %v\n", space.Pages)
		}
	}

	pages, err := pageRepo.FindAll(context.TODO())
	if err == nil {
		for _, page := range pages {
			fmt.Printf(" - %v\n", page)
		}
	}

	bookmarks, err := bookmarkRepo.FindAll(context.TODO())
	if err == nil {
		for _, bookmark := range bookmarks {
			fmt.Printf(" - %v\n", bookmark)
		}
	}

	// Feature Service
	featureService := service.NewFeatureService(spaceRepo, pageRepo, bookmarkRepo)
	spaceById, err := featureService.GetSpaceById(context.TODO(), space.ID)
	if err != nil {
		return
	}
	fmt.Printf(" - %v\n", spaceById)
}
