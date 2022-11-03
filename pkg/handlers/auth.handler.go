package handlers

import (
	"strings"
	"time"

	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = utils.GetEnv("JWT_SECRET_KEY", "s3cr3t")
var jwtExpirationHours = utils.MustParseInt64(utils.GetEnv("JWT_EXPIRE_HOURS", "72"))

type LoginData struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type Registration struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Email    string `json:"email" xml:"email" form:"email"`
}

func Login(c *fiber.Ctx) error {
	var payload LoginData
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if payload.Username == "" || payload.Password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	username := strings.ToLower(payload.Username)

	var user models.User
	result := database.GetConnection().Where("lower(username) = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return c.SendStatus(fiber.StatusForbidden)
	}

	claims := jwt.MapClaims{
		"sub":    username,
		"userid": user.ID,
		"exp":    time.Now().Add(time.Hour * time.Duration(jwtExpirationHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString(t)
}

func Register(c *fiber.Ctx) error {
	var payload Registration
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if payload.Username == "" || payload.Password == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	username := strings.ToLower(payload.Username)

	var user models.User
	result := database.GetConnection().Where("lower(username) = ?", username).Find(&user)

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).SendString("User already exists")
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	user = models.User{
		Username: username,
		Password: hashedPassword,
		Email:    payload.Email,
	}
	database.GetConnection().Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}
