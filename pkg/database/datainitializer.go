package database

import (
	"context"
	"fmt"
	"os"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/teris-io/shortid"
)

const (
	admin        = "admin"
	passwordFile = "adminpassword"
)

func InitializeAdmin(c context.Context, q *gen.Queries) {
	e, err := q.CountUserByName(c, admin)
	if err != nil || e > 0 {
		fmt.Println("Admin already exists")
		return
	}

	shortId := shortid.MustGenerate()
	pwd, err := utils.HashPassword(shortId)
	if err != nil {
		panic(err.Error())
	}

	id, err := q.CreateUser(c, gen.CreateUserParams{
		Username: admin,
		Password: pwd,
		Email:    "admin@admin",
	})
	if err != nil {
		panic(err.Error())
	}
	n, err := q.CountUserRole(c, gen.CountUserRoleParams{
		UserID: id,
		Name:   admin,
	})
	if err != nil {
		panic(err.Error())
	}
	if n < 1 {
		role, err := q.FindRoleByName(c, admin)
		if err != nil {
			panic(err.Error())
		}
		q.InsertUserRole(c, gen.InsertUserRoleParams{
			UserID: id,
			RoleID: role.ID,
		})
	}

	file, err := os.Create("adminpassword")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("File created admin password file")
	defer file.Close()

	_, err = file.WriteString(shortId)
	if err != nil {
		panic(err.Error())
	}
}
