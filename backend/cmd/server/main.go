package main

import (
	"backend/cmd/server/handlers"
	"backend/internal/db"
	"backend/internal/events"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	dispatcher := events.NewEventDispatcher()

	pageRepo := repository.NewPageRepository(database.Collection("pages"), dispatcher)
	spaceRepo := repository.NewSpaceRepository(database.Collection("spaces"), dispatcher)
	bookmarkRepo := repository.NewBookmarkRepository(database.Collection("bookmarks"), dispatcher)

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
	featureService := service.NewFeatureService(spaceRepo, pageRepo)
	spaceById, err := featureService.GetSpaceById(context.TODO(), space.ID)
	if err != nil {
		return
	}
	fmt.Printf(" - %v\n", spaceById)
}
