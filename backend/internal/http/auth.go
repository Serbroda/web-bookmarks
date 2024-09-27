package http

import (
	"backend/internal"
	"backend/internal/product"
	security2 "backend/internal/security"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	RefreshToken security2.Jwt `json:"refreshToken"`
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
	UserService *product.UserServiceImpl
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
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := security2.HashBcrypt(payload.Password)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	entity := &internal.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	if payload.Username != nil && *payload.Username != "" {
		entity.Username = *payload.Username
	}

	err = si.UserService.Create(entity)

	if err != nil {
		if errors.Is(err, product.ErrUsernameAlreadyExists) {
			return ctx.String(http.StatusConflict, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (si *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	entity, err := si.UserService.GetUserByEmailOrUsername(payload.User)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.String(http.StatusUnauthorized, "bad login credentials")
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if !security2.CheckBcryptHash(payload.Password, entity.Password) {
		return ctx.String(http.StatusUnauthorized, "bad login credentials")
	}

	tokenPair, err := security2.GenerateJwtPair(entity)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, tokenPair)
}

func (si *AuthHandler) RefreshToken(ctx echo.Context) error {
	var payload RefreshTokenRequest
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := payload.Validate(); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	token, err := security2.ParseJwt(payload.RefreshToken)

	if err != nil {
		return middleware.ErrJWTInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return middleware.ErrJWTInvalid
	}

	sub := claims["sub"].(string)
	entity, err := si.UserService.GetById(sub)

	if err != nil {
		return echo.ErrUnauthorized
	}

	tokenPair, err := security2.GenerateJwtPair(entity)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, tokenPair)
}
