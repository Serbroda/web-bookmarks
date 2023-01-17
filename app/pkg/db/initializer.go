package db

import (
	"context"
	"fmt"
	"github.com/Serbroda/ragbag/app/pkg/services"
	"github.com/Serbroda/ragbag/app/sqlc"
	"github.com/teris-io/shortid"
	"log"
	"os"
)

const (
	admin        = "admin"
	passwordFile = "adminpassword"
)

func InitializeAdmin(c context.Context, s *services.Services) {
	if s.ExistsUser(c, admin) {
		return
	}

	log.Println("initializing admin user")
	shortId := shortid.MustGenerate()

	user, err := s.CreateUserWithRoles(c, sqlc.CreateUserParams{
		Username:  admin,
		Password:  shortId,
		Email:     "admin@admin",
		Active:    true,
		FirstName: "Admin",
		LastName:  "Admin",
	}, []string{"ADMIN"})

	if err != nil {
		panic(err.Error())
	}

	_, err = s.CreateSpace(c, sqlc.CreateSpaceParams{
		ShortID: shortid.MustGenerate(),
		OwnerID: user.ID,
		Name:    "Admin's Space",
	})

	if err != nil {
		panic(err.Error())
	}

	file, err := os.Create(passwordFile)
	if err != nil {
		panic(err.Error())
	}
	//abs, err := filepath.Abs(passwordFile)
	if err == nil {
		//log.Printf("adminpassword file created. Execute 'cat %s' to view initial password.", abs)
		log.Printf("adminpassword file created: %s", shortId)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n", shortId))
	if err != nil {
		panic(err.Error())
	}
}
