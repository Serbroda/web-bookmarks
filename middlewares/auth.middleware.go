package middlewares

import (
	"webcrate/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = utils.GetEnv("JWT_SECRET_KEY", "s3cr3t")

type Authentication struct {
	Id      uint
	Subject string
}

func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecretKey),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			if user == nil {
				return c.Next()
			}
			claims := user.Claims.(jwt.MapClaims)

			id := claims["userid"].(float64)
			sub := claims["sub"].(string)

			c.Locals("authentication", Authentication{
				Id:      uint(id),
				Subject: sub,
			})
			return c.Next()
		},
	})
}
