package handlers

import (
	"backend/internal/model"
	"backend/internal/security"
	"backend/internal/service"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RefreshTokenRequest struct {
	RefreshToken security.Jwt `json:"refresh_token"`
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
	if err != nil || payload.Email == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	hashedPassword, err := security.HashBcrypt(payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	user := &model.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if payload.Username != "" {
		user.Username = payload.Username
	}

	err = si.UserService.CreateUser(user)

	if err != nil {
		if errors.Is(err, service.ErrUsernameAlreadyExists) {
			return ctx.String(http.StatusConflict, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, user)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.User == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := si.UserService.GetUserByEmailOrUsername(payload.User)
	if err != nil || !security.CheckBcryptHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	tokenPair, err := security.GenerateJwtPair(user)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	var payload RefreshTokenRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.RefreshToken == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	token, err := security.ParseJwt(payload.RefreshToken)

	if err != nil {
		return middleware.ErrJWTInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return middleware.ErrJWTInvalid
	}

	sub := claims["sub"].(string)
	user, err := si.UserService.GetUserById(sub)

	if err != nil {
		return echo.ErrUnauthorized
	}

	tokenPair, err := security.GenerateJwtPair(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}

	return ctx.JSON(http.StatusOK, tokenPair)
}
