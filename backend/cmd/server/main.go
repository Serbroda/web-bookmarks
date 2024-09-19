package main

import (
	"backend/cmd/server/handlers"
	"backend/internal/db"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/security"
	"backend/internal/service"
	"context"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	_, database := db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	userRepo := repository.NewUserRepository(database.Collection("users"))
	pageRepo := repository.NewPageRepository(database.Collection("pages"))
	spaceRepo := repository.NewSpaceRepository(database.Collection("spaces"))
	bookmarkRepo := repository.NewBookmarkRepository(database.Collection("bookmarks"))

	userService := service.NewUserService(userRepo)
	contentService := service.NewContentService(spaceRepo, pageRepo, bookmarkRepo)

	tryDB(spaceRepo)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	handlers.RegisterUsersHandlers(api, handlers.UsersHandler{UserService: userService}, "/v1")
	handlers.RegisterContentHandlers(api, handlers.ContentHandler{ContentService: contentService}, "/v1")

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func tryDB(spaceRepo *repository.SpaceRepository) {
	space := &model.Space{
		Name:   "test struct",
		Shared: make([]model.UserIdWithRole, 0),
	}

	userId, err := bson.ObjectIDFromHex("66eb41a0829447497723b259")
	if err != nil {
		panic(err)
	}

	space.Shared = append(space.Shared, model.UserIdWithRole{
		UserID: userId,
		Role:   "OWNER",
	})

	err = spaceRepo.Save(context.TODO(), space)
	if err != nil {
		panic(err)
	}

	founds, err := spaceRepo.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	for _, found := range founds {
		fmt.Printf(" - %v\n", found)
	}
}
