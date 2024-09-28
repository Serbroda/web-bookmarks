package http

import (
	"backend/internal/dto"
	"backend/internal/security"
	"backend/internal/services"
	"backend/internal/sqlc"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"strconv"
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
	UserService *services.UserServiceImpl
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	if err := BindAndValidate(ctx, &payload); err != nil {
		return err
	}

	hashedPassword, err := security.HashBcrypt(payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	params := sqlc.CreateParams{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if payload.Username != nil && *payload.Username != "" {
		params.Username = *payload.Username
	}

	user, err := si.UserService.Create(params)

	if err != nil {
		if errors.Is(err, services.ErrUsernameAlreadyExists) {
			return ctx.String(http.StatusConflict, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, copier.Copy(&user, dto.UserDto{}))
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	if err := BindAndValidate(ctx, &payload); err != nil {
		return err
	}

	entity, err := si.UserService.GetByEmailOrUsername(payload.User)
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
	if err := BindAndValidate(ctx, &payload); err != nil {
		return err
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
	userId, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return err
	}
	entity, err := si.UserService.GetById(userId)

	if err != nil {
		return echo.ErrUnauthorized
	}

	tokenPair, err := security.GenerateJwtPair(entity)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, tokenPair)
}
