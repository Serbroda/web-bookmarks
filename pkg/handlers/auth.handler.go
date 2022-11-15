package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtSecretKey = utils.GetEnv("JWT_SECRET_KEY", "s3cr3t")
var jwtExpirationHours = utils.MustParseInt64(utils.GetEnv("JWT_EXPIRE_HOURS", "72"))

type JwtCustomClaims struct {
	Subject string `json:"sub"`
	UserId  uint   `json:"userid"`
	jwt.StandardClaims
}

type LoginData struct {
	Username string `param:"username" query:"username" json:"username" xml:"username" form:"username"`
	Password string `param:"password" query:"password" json:"password" xml:"password" form:"password"`
}

type Registration struct {
	Username string `param:"username" query:"username" json:"username" xml:"username" form:"username"`
	Password string `param:"password" query:"password" json:"password" xml:"password" form:"password"`
	Email    string `param:"email" query:"email" json:"email" xml:"email" form:"email"`
}

func Login(c echo.Context) error {
	var payload LoginData
	err := c.Bind(&payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if payload.Username == "" || payload.Password == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}

	username := strings.ToLower(payload.Username)

	var user models.User
	result := database.GetConnection().Where("lower(username) = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "User not found")
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return c.String(http.StatusForbidden, "Wrong email or password")
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
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusOK, t)
}

func Register(c echo.Context) error {
	var payload Registration
	err := c.Bind(&payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if payload.Username == "" || payload.Password == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}

	username := strings.ToLower(payload.Username)

	var user models.User
	result := database.GetConnection().Where("lower(username) = ?", username).Find(&user)

	if result.RowsAffected > 0 {
		return c.String(http.StatusConflict, "User already exists")
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	user = models.User{
		Username: username,
		Password: hashedPassword,
		Email:    payload.Email,
	}
	database.GetConnection().Create(&user)

	return c.JSON(http.StatusCreated, user)
}
