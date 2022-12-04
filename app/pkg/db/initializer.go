package db

import (
	"context"
)

const (
	admin        = "admin"
	passwordFile = "adminpassword"
)

func InitializeAdmin(c context.Context) {
	/*if s.ExistsUser(c, admin) {
		return
	}

	fmt.Println("initializing admin user")
	shortId := shortid.MustGenerate()

	_, err := s.CreateUserWithRoles(c, gen.CreateUserParams{
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

	file, err := os.Create("adminpassword")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("adminpassword file created with initial password: %s\n", shortId)
	defer file.Close()

	_, err = file.WriteString(shortId)
	if err != nil {
		panic(err.Error())
	}*/
}
