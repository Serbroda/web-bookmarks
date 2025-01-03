package handlers

import (
	"context"
	"errors"
	"github.com/Serbroda/bookmark-manager/internal/models"
	"github.com/Serbroda/bookmark-manager/internal/security"
	"github.com/Serbroda/bookmark-manager/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

type LoginRequest struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationRequest struct {
	Email    string  `json:"email" validate:"email,required"`
	Password string  `json:"password" validate:"required"`
	Username *string `json:"username,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken security.Jwt `json:"refreshToken" validate:"required"`
}

type AuthHandler struct {
	UserService services.UserService
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	hashedPassword, err := security.HashBcrypt(payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	entity := &models.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if payload.Username != nil && *payload.Username != "" {
		entity.Username = *payload.Username
	}

	_, err = si.UserService.CreateUser(context.TODO(), entity)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	entity, err := si.UserService.GetByEmailOrUsername(context.TODO(), payload.User)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.String(http.StatusUnauthorized, "bad login credentials")
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if !security.CheckBcryptHash(payload.Password, entity.Password) {
		return ctx.String(http.StatusUnauthorized, "bad login credentials")
	}

	tokenPair, err := security.GenerateJwtPair(entity)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	var payload RefreshTokenRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	token, err := security.ParseJwt(payload.RefreshToken)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	sub := claims["sub"].(string)
	entity, err := si.UserService.GetUserById(context.TODO(), sub)

	if err != nil {
		return echo.ErrUnauthorized
	}

	tokenPair, err := security.GenerateJwtPair(entity)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, tokenPair)
}
