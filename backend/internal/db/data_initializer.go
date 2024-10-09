package db

import (
	"backend/internal/common/random"
	"backend/internal/security"
	"backend/internal/sqlc"
	"context"
	"log"
)

const (
	adminEmail       = "admin@admin.net"
	adminUsername    = "amin"
	adminDisplayname = "Admin"
)

func InitializeData(queries *sqlc.Queries) {
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

	_, err = queries.CreateUser(context.TODO(), sqlc.CreateUserParams{
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
