package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/gen/public"
	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtSecretKey = utils.GetEnvFallback("JWT_SECRET_KEY", "s3cr3t")
var jwtExpirationHours = utils.MustParseInt64(utils.GetEnvFallback("JWT_EXPIRE_HOURS", "72"))

type PublicServerInterfaceImpl struct {
}

type JwtCustomClaims struct {
	Subject string `json:"sub"`
	UserId  int64  `json:"userid"`
	jwt.StandardClaims
}

func (si *PublicServerInterfaceImpl) Login(ctx echo.Context) error {
	var payload public.LoginDto
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if payload.Username == nil || payload.Password == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	username := strings.ToLower(*payload.Username)

	user, err := database.Queries.GetUserByUsername(ctx.Request().Context(), *payload.Username)
	if err == nil {
		return ctx.String(http.StatusNotFound, "User not found")
	}

	if !utils.CheckPasswordHash(*payload.Password, user.Password) {
		return ctx.String(http.StatusForbidden, "Wrong email or password")
	}

	claims := &JwtCustomClaims{
		Subject: username,
		UserId:  user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(jwtExpirationHours)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.String(http.StatusOK, t)
}

func (si *PublicServerInterfaceImpl) Register(ctx echo.Context) error {
	var payload public.RegistrationDto
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if payload.Username == nil || payload.Password == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	username := strings.ToLower(*payload.Username)

	user, err := database.Queries.GetUserByUsername(ctx.Request().Context(), *payload.Username)
	if err == nil {
		return ctx.String(http.StatusNotFound, "User not found")
	}

	hashedPassword, _ := utils.HashPassword(*payload.Password)

	params := gen.CreateUserParams{
		Username: username,
		Password: hashedPassword,
		Email:    *payload.Email,
	}

	res, err := database.Queries.CreateUser(ctx.Request().Context(), params)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	user, err = database.Queries.GetUser(ctx.Request().Context(), id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, user)
}
