package handlers

import (
	"backend/models"
	"backend/security"
	"backend/services"
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

func (j *LoginRequest) Validate() *ConstraintViolationError {
	var violations ConstraintViolationError

	if len(j.User) == 0 {
		violations.AddViolation("user", "user must be set")
	}
	if len(j.Password) == 0 {
		violations.AddViolation("password", "password must be set")
	}

	if len(violations.Violations) > 0 {
		return &violations
	}
	return nil
}

type RegistrationRequest struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Username *string `json:"username,omitempty"`
}

func (j *RegistrationRequest) Validate() *ConstraintViolationError {
	var violations ConstraintViolationError

	if len(j.Email) == 0 {
		violations.AddViolation("user", "user must be set")
	}
	if len(j.Password) == 0 {
		violations.AddViolation("password", "password must be set")
	}

	if len(violations.Violations) > 0 {
		return &violations
	}
	return nil
}

type RefreshTokenRequest struct {
	RefreshToken security.Jwt `json:"refreshToken"`
}

func (j *RefreshTokenRequest) Validate() *ConstraintViolationError {
	var violations ConstraintViolationError

	if len(j.RefreshToken) == 0 {
		violations.AddViolation("refreshToken", "refreshToken must be set")
	}

	if len(violations.Violations) > 0 {
		return &violations
	}
	return nil
}

type AuthHandler struct {
	UserService *services.UserService
}

func RegisterAuthHandlers(e *echo.Echo, c AuthHandler, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.POST(baseUrl+"/auth/signup", c.Register, middlewares...)
	e.POST(baseUrl+"/auth/login", c.Login, middlewares...)
	e.POST(baseUrl+"/auth/refresh_token", c.RefreshToken, middlewares...)
}

func (si *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	if err := payload.Validate(); err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	hashedPassword, err := security.HashBcrypt(payload.Password)
	if err != nil {
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	user := &models.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if payload.Username != nil && *payload.Username != "" {
		user.Username = *payload.Username
	}

	err = si.UserService.CreateUser(user)

	if err != nil {
		if errors.Is(err, services.ErrUsernameAlreadyExists) {
			return handleError(ctx, err, http.StatusConflict)
		} else {
			return handleError(ctx, err, http.StatusInternalServerError)
		}
	}

	return ctx.JSON(http.StatusOK, user)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	if err := payload.Validate(); err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	user, err := si.UserService.GetUserByEmailOrUsername(payload.User)
	if err != nil || !security.CheckBcryptHash(payload.Password, user.Password) {
		return handleError(ctx, err, http.StatusUnauthorized)
	}

	tokenPair, err := security.GenerateJwtPair(user)

	if err != nil {
		return handleError(ctx, err, http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	var payload RefreshTokenRequest
	err := ctx.Bind(&payload)
	if err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
	}

	if err := payload.Validate(); err != nil {
		return handleError(ctx, err, http.StatusBadRequest)
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
		return handleError(ctx, err, http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, tokenPair)
}
