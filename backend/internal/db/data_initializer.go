package db

import (
	"backend/internal/common/random"
	sqlc2 "backend/internal/db/sqlc"
	"backend/internal/security"
	"context"
	"log"
)

const (
	adminEmail       = "admin@admin.net"
	adminUsername    = "amin"
	adminDisplayname = "Admin"
)

func InitializeData(queries *sqlc2.Queries) {
	count, err := queries.CountUserByUsername(context.TODO(), adminUsername)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return
	}

	passwordPlain := random.RandomStringWithCharset(12, random.CharsetAlphaNumeric)
	password, err := security.HashBcrypt(passwordPlain)
	if err != nil {
		return
	}

	_, err = queries.CreateUser(context.TODO(), sqlc2.CreateUserParams{
		Email:       adminEmail,
		Username:    adminUsername,
		Password:    password,
		DisplayName: toPointer(adminDisplayname),
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Initialized %s with password: %s\n", adminUsername, passwordPlain)
}

func toPointer(s string) *string {
	return &s
}
