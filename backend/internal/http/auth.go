package http

import (
	"backend/internal/db/sqlc"
	"backend/internal/dtos"
	"backend/internal/security"
	"backend/internal/services"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type LoginRequest struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationRequest struct {
	Email       string  `json:"email" validate:"email,required"`
	Password    string  `json:"password" validate:"required"`
	Username    *string `json:"username" validate:"required"`
	DisplayName *string `json:"displayName,omitempty"`
}

type AuthHandler struct {
	UserService *services.UserServiceImpl
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh", c.Refresh, middlewares...)
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

	params := sqlc.CreateUserParams{
		Email:       payload.Email,
		Password:    hashedPassword,
		Username:    *payload.Username,
		DisplayName: payload.DisplayName,
	}

	user, err := si.UserService.CreateUser(params)

	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyExists) || errors.Is(err, services.ErrUsernameAlreadyExists) {
			return ctx.String(http.StatusConflict, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, dtos.UserDtoFromUser(user))
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	if err := BindAndValidate(ctx, &payload); err != nil {
		return err
	}

	entity, err := si.UserService.GetUserByEmailOrUsername(payload.User)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    *tokenPair.RefreshToken,
		Expires:  tokenPair.RefreshTokenExpiration, // 7 days
		HttpOnly: true,
		Secure:   true, // Set to true in production (HTTPS only)
	})
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) Refresh(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "refresh token missing",
		})
	}

	token, err := security.VerifyRefreshToken(cookie.Value)
	if err != nil {
		return middleware.ErrJWTInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return middleware.ErrJWTInvalid
	}

	sub := claims["sub"].(float64)
	userId := int64(sub)
	entity, err := si.UserService.GetUserById(userId)

	if err != nil {
		return echo.ErrUnauthorized
	}

	tokenPair, err := security.GenerateJwtPair(entity)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    *tokenPair.RefreshToken,
		Expires:  tokenPair.RefreshTokenExpiration, // 7 days
		HttpOnly: true,
		Secure:   true, // Set to true in production (HTTPS only)
	})

	return ctx.JSON(http.StatusOK, tokenPair)
}
