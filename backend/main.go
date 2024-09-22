package main

import (
	"backend/db"
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/security"
	"backend/services"
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

	userRepo := repositories.NewUserRepository(database.Collection("users"))
	pageRepo := repositories.NewPageRepository(database.Collection("pages"))
	spaceRepo := repositories.NewSpaceRepository(database.Collection("spaces"))
	bookmarkRepo := repositories.NewBookmarkRepository(database.Collection("bookmarks"))

	userService := services.NewUserService(userRepo)
	contentService := services.NewContentService(spaceRepo, pageRepo, bookmarkRepo)

	tryDB(spaceRepo, pageRepo)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlers.RegisterAuthHandlers(e, handlers.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	handlers.RegisterUsersHandlers(api, handlers.UsersHandler{UserService: userService}, "/v1")
	handlers.RegisterSpaceHandlers(api, handlers.SpaceHandler{ContentService: contentService}, "/v1")

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func tryDB(spaceRepo *repositories.SpaceRepository, pageRepo *repositories.PageRepository) {
	space := &models.Space{
		Name:   "Space 1",
		Shared: make([]models.UserIdWithRole, 0),
	}

	userId, err := bson.ObjectIDFromHex("66f07f396e64446f862e37da")
	if err != nil {
		panic(err)
	}

	space.Shared = append(space.Shared, models.UserIdWithRole{
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

	page0 := &models.Page{
		Name:    "Page R0",
		SpaceID: space.ID,
	}
	err = pageRepo.Save(context.TODO(), page0)
	if err != nil {
		panic(err)
	}
	page1 := &models.Page{
		Name:    "Page R1",
		SpaceID: space.ID,
	}
	err = pageRepo.Save(context.TODO(), page1)
	if err != nil {
		panic(err)
	}
	page11 := &models.Page{
		Name:         "Page R21.1",
		SpaceID:      space.ID,
		ParentPageID: &page1.ID,
	}
	err = pageRepo.Save(context.TODO(), page11)
	if err != nil {
		panic(err)
	}
	page12 := &models.Page{
		Name:         "Page R21.2",
		SpaceID:      space.ID,
		ParentPageID: &page1.ID,
	}
	err = pageRepo.Save(context.TODO(), page12)
	if err != nil {
		panic(err)
	}
	page121 := &models.Page{
		Name:         "Page R21.2.1",
		SpaceID:      space.ID,
		ParentPageID: &page12.ID,
	}
	err = pageRepo.Save(context.TODO(), page121)
	if err != nil {
		panic(err)
	}
}
