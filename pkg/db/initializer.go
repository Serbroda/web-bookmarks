package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/teris-io/shortid"
)

const (
	admin        = "admin"
	passwordFile = "adminpassword"
)

func InitializeAdmin(c context.Context, s *services.Service) {
	if s.UserService.ExistsUser(c, admin) {
		return
	}

	fmt.Println("initializing admin user")
	shortId := shortid.MustGenerate()

	_, err := s.UserService.CreateUserWithRoles(c, gen.CreateUserParams{
		Username:           admin,
		Password:           shortId,
		Email:              "admin@admin",
		Active:             true,
		MustChangePassword: true,
	}, []string{"ADMIN"})

	if err != nil {
		panic(err.Error())
	}

	file, err := os.Create("adminpassword")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("adminpassword file created with initial password: %s\n", shortId)
	defer file.Close()

	_, err = file.WriteString(shortId)
	if err != nil {
		panic(err.Error())
	}
}
