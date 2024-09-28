package main

import (
	"backend/internal"
	"backend/internal/http"
	"backend/internal/mongodb"
	"backend/internal/product"
	"backend/internal/security"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	_, database := mongodb.Connect("mongodb://localhost:27017")
	defer mongodb.CloseConnection()

	userRepo := mongodb.NewUserRepository(database.Collection("users"))
	pageRepo := mongodb.NewPageRepository(database.Collection("pages"))
	spaceRepo := mongodb.NewSpaceRepository(database.Collection("spaces"))
	bookmarkRepo := mongodb.NewBookmarkRepository(database.Collection("bookmarks"))

	userService := product.NewUserService(userRepo)
	contentService := product.NewContentService(spaceRepo, pageRepo, bookmarkRepo)

	tryDB(spaceRepo, pageRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	http.RegisterAuthHandlers(e, http.AuthHandler{
		UserService: userService,
	}, "")

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(security.CreateJwtConfig()))
	http.RegisterUsersHandlers(api, http.UsersHandler{}, "/v1")
	http.RegisterSpaceHandlers(api, http.SpaceHandler{ContentService: contentService}, "/v1")

	printRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func printRoutes(e *echo.Echo) {
	fmt.Printf("Registered following routes\n\n")
	for _, r := range e.Routes() {
		fmt.Printf(" - %v %v\n", r.Method, r.Path)
	}
}

func tryDB(spaceRepo *mongodb.SpaceRepository, pageRepo *mongodb.PageRepository) {
	space := &internal.Space{
		Name:   "Space 1",
		Shared: make([]internal.UserIdWithRole, 0),
	}

	userId, err := bson.ObjectIDFromHex("66f07f396e64446f862e37da")
	if err != nil {
		panic(err)
	}

	space.Shared = append(space.Shared, internal.UserIdWithRole{
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

	page0 := &internal.Page{
		Name:    "Page R0",
		SpaceID: space.ID,
	}
	err = pageRepo.Save(context.TODO(), page0)
	if err != nil {
		panic(err)
	}
	page1 := &internal.Page{
		Name:    "Page R1",
		SpaceID: space.ID,
	}
	err = pageRepo.Save(context.TODO(), page1)
	if err != nil {
		panic(err)
	}
	page11 := &internal.Page{
		Name:         "Page R21.1",
		SpaceID:      space.ID,
		ParentPageID: &page1.ID,
	}
	err = pageRepo.Save(context.TODO(), page11)
	if err != nil {
		panic(err)
	}
	page12 := &internal.Page{
		Name:         "Page R21.2",
		SpaceID:      space.ID,
		ParentPageID: &page1.ID,
	}
	err = pageRepo.Save(context.TODO(), page12)
	if err != nil {
		panic(err)
	}
	page121 := &internal.Page{
		Name:         "Page R21.2.1",
		SpaceID:      space.ID,
		ParentPageID: &page12.ID,
	}
	err = pageRepo.Save(context.TODO(), page121)
	if err != nil {
		panic(err)
	}
}
