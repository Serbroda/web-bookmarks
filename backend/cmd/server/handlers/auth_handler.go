package handlers

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	UserService *service.UserService
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	hashedPassword, err := utils.HashBcrypt(payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	user := &model.User{
		Username: payload.Username,
		Password: hashedPassword,
		Email:    payload.Email,
	}

	err = si.UserService.CreateUser(user)

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			return ctx.String(http.StatusConflict, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusInternalServerError, user)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := si.UserService.GetUserByUsername(payload.Username)
	if err != nil || !utils.CheckBcryptHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	tokenPair, err := utils.GenerateJwtPair(user)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	return ctx.String(http.StatusInternalServerError, "Not Implemented")
}
