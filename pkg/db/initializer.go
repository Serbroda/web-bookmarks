package db

import (
	"context"
	"fmt"
	"github.com/Serbroda/ragbag/pkg/security"
	"log"

	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/sqlc"
)

var (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

type DataInitializer struct {
	Queries     *sqlc.Queries
	UserService services.UserService
}

func (di *DataInitializer) Initialize() {
	di.initializeRoles()
	di.initializeAdminUser()
}

func (di *DataInitializer) initializeRoles() {
	di.createRoleIfNotExist(RoleAdmin)
	di.createRoleIfNotExist(RoleUser)
}

func (di *DataInitializer) initializeAdminUser() {
	_, err := di.UserService.FindByUsername(context.Background(), "admin")
	if err == nil {
		return
	}
	roles, err := di.Queries.FindRolesByNames(context.Background(), []string{RoleAdmin})
	if err != nil {
		panic("failed to create admin")
	}
	password := security.RandomString(8)
	passwordHashed, err := security.HashBcrypt(password)
	if err != nil {
		log.Fatal("failed to initialize admin user")
	}
	fmt.Printf("admin generated with password: %s\n", password)
	admin, err := di.UserService.Create(context.Background(), services.CreateUser{
		CreateUserParams: sqlc.CreateUserParams{
			Username: "admin",
			Password: passwordHashed,
			Email:    "admin@example.com",
			Active:   true,
		},
		Roles: roles,
	})
	if err != nil {
		panic("failed to create admin")
	}
	fmt.Printf("Created admin user with id: %v\n", admin.ID)
}

func (di *DataInitializer) createRoleIfNotExist(name string) {
	_, err := di.Queries.FindRoleByName(context.Background(), name)
	if err == nil {
		return
	}
	id, err := di.Queries.CreateRole(context.Background(), sqlc.CreateRoleParams{
		Name: name,
	})
	if err != nil {
		panic("failed to create role")
	}
	fmt.Printf("Created role '%v' with id: %v\n", name, id)
}
