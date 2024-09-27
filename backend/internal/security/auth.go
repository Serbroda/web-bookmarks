package security

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

var (
	ContextKeyAuthentication = "authentication"
)

type Authentication struct {
	Subject string
	UserId  bson.ObjectID
}

func CreateJwtConfig() echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(JwtSecretKey),
		ContextKey: "token",
		SuccessHandler: func(c echo.Context) {
			token, ok := c.Get("token").(*jwt.Token)
			if !ok {
				return
			}
			auth, err := ParseToken(token)
			if err != nil {
				return
			}
			c.Set(ContextKeyAuthentication, auth)
		},
	}
}

func ParseToken(token *jwt.Token) (Authentication, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Authentication{}, errors.New("failed to get claims of token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return Authentication{}, errors.New("failed to get sub from claims")
	}

	userId, err := bson.ObjectIDFromHex(sub)
	if err != nil {
		return Authentication{}, errors.New("failed to create ObjectID from sub")
	}

	return Authentication{
		Subject: sub,
		UserId:  userId,
	}, nil
}

func GetAuthentication(ctx echo.Context) (Authentication, error) {
	auth, ok := ctx.Get(ContextKeyAuthentication).(Authentication)
	if !ok {
		return auth, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	return auth, nil
}
