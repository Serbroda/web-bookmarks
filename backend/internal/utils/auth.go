package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	ContextKeyAuthentication = "authentication"
	RoleAdmin                = "ADMIN"
)

type Authentication struct {
	Subject string
	Roles   []string
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
	/*userId, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return Authentication{}, errors.New("failed parse int of sub")
	}
	roleInterfaces, ok := claims["roles"].([]interface{})
	if !ok {
		return Authentication{}, errors.New("failed to get roles from claims")
	}
	roles := []string{}
	for _, ri := range roleInterfaces {
		roles = append(roles, ri.(string))
	}*/
	return Authentication{
		Subject: sub,
		//Roles:   roles,
	}, nil
}

func HasAnyRoleMiddleware(roles ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authentication, ok := c.Get(ContextKeyAuthentication).(Authentication)
			if !ok {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}
			if !IncludesAnyRole(authentication.Roles, roles...) {
				return c.String(http.StatusForbidden, "Forbidden")
			}
			return next(c)
		}
	}
}

func IncludesAnyRole(authRoles []string, roles ...string) bool {
	for _, ur := range authRoles {
		for _, r := range roles {
			if strings.EqualFold(ur, r) {
				return true
			}
		}
	}
	return false
}
