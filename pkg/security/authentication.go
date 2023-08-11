package security

import (
	"net/http"
	"strconv"

	"github.com/Serbroda/ragbag/pkg/user"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	ContextKeyAuthentication = "authentication"
)

type Authentication struct {
	Token *jwt.Token
	User  *user.User
}

func CreateJwtConfig(userService user.UserService) echojwt.Config {
	return echojwt.Config{
		SigningKey: []byte(JwtSecretKey),
		ContextKey: "token",
		SuccessHandler: func(c echo.Context) {
			token, ok := c.Get("token").(*jwt.Token)
			if !ok {
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return
			}
			sub := claims["sub"].(string)
			id, err := strconv.ParseInt(sub, 10, 64)
			if err != nil {
				return
			}
			user, err := userService.FindOne(c.Request().Context(), id)
			if err == nil {
				c.Set(ContextKeyAuthentication, Authentication{
					Token: token,
					User:  &user,
				})
			}
		},
	}
}

func HasAnyRoleMiddleware(roles ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u := c.Get(ContextKeyAuthentication)
			authentication, ok := u.(Authentication)
			if !ok || authentication.User == nil {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}
			if !authentication.User.HasAnyRole(roles...) {
				return c.String(http.StatusForbidden, "Forbidden")
			}
			return next(c)
		}
	}
}
