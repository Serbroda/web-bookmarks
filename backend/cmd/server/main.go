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
)

// https://golang.withcodeexample.com/blog/top-databases-with-golang-in-2024/#:~:text=for%20more%20examples-,MongoDB,or%20Not%20only%20SQL%20database.
func main() {
	db.Connect("mongodb://localhost:27017")
	defer db.CloseConnection()

	spaceRepo := repository.NewMongoSpaceRepository(db.Database.Collection("spaces"))
	space := &model.Space{
		Name: "First space",
	}
	id, err := spaceRepo.Save(context.TODO(), space)
	if err != nil {
		fmt.Println(err)
	}

	pageRepo := repository.NewMongoPageRepository(db.Database.Collection("pages"))
	page := model.Page{
		Name:    "First page",
		SpaceID: space.ID,
	}
	id, err = pageRepo.Save(context.TODO(), &page)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	foundSpace, err := spaceRepo.FindByID(context.TODO(), space.ID.Hex())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(foundSpace)
	}

	id, err = spaceRepo.Save(context.TODO(), foundSpace)
	if err != nil {
		fmt.Println(err)
	}

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
