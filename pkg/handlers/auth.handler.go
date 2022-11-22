package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Serbroda/ragbag/gen"
	"github.com/Serbroda/ragbag/gen/public"
	"github.com/Serbroda/ragbag/pkg/db"
	"github.com/Serbroda/ragbag/pkg/services"
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
	if err != nil || payload.Username == nil || payload.Password == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	username := strings.ToLower(*payload.Username)

	user, err := db.Queries.FindUserByName(ctx.Request().Context(), *payload.Username)
	if err != nil || user.ID < 1 || !utils.CheckPasswordHash(*payload.Password, user.Password) {
		return ctx.String(http.StatusNotFound, "invalid login")
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
	if err != nil || payload.Username == nil || payload.Password == nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if services.Services.UserService.ExistsUser(ctx.Request().Context(), *payload.Password) {
		return ctx.String(http.StatusConflict, "user already exists")
	}

	hashedPassword, _ := utils.HashPassword(*payload.Password)

	user, err := services.Services.UserService.CreateUser(ctx.Request().Context(), gen.CreateUserParams{
		Username: strings.ToLower(*payload.Username),
		Password: hashedPassword,
		Email:    *payload.Email,
	})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, &public.UserDto{
		Id:       &user.ID,
		Username: &user.Username,
	})
}
