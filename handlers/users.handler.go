package handlers

import (
	"net/http"
	"strconv"
	"webcrate/database"
	"webcrate/middlewares"
	"webcrate/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.GetConnection().Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userId, _ := strconv.ParseUint(c.Params("userId"), 10, 32)

	var user models.User
	result := database.GetConnection().Find(&user, userId)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func GetMe(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var user models.User
	result := database.GetConnection().Find(&user, authentication.Id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	database.GetConnection().Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}
