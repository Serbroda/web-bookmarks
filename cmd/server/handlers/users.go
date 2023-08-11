package handlers

import (
	"net/http"

	"github.com/Serbroda/ragbag/pkg/security"
	"github.com/Serbroda/ragbag/pkg/user"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UserService user.UserService
}

func RegisterUsersHandlers(e *echo.Echo, h UsersHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/users", h.GetUsers, append(middlewares, security.HasAnyRoleMiddleware("ADMIN"))...)
}

func (h *UsersHandler) GetUsers(ctx echo.Context) error {
	auth, ok := ctx.Get(security.ContextKeyAuthentication).(security.Authentication)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Unauthorized")
	}
	/*token, ok := ctx.Get("user").(*jwt.Token)
	fmt.Println(token)
	if !ok {
		return ctx.String(http.StatusUnauthorized, "Missing token")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	sub := claims["sub"].(string)
	id, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return err
	}
	fmt.Printf("User context id = %v", id)*/
	/*u := ctx.Get("user")
	res, ok := u.(user.User)
	if !ok {
		return errors.New("no user_id")
	}
	fmt.Printf("user=%v", res)*/
	return ctx.JSON(http.StatusOK, auth)
}
