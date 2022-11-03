package handlers

import (
	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/middlewares"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ChangePasswordDto struct {
	OldPassword string `json:"oldPassword" xml:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" xml:"newPassword" form:"newPassword"`
}

func GetMe(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)

	user, err := services.FindUserById(authentication.Id)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func ChangePassword(c *fiber.Ctx) error {
	var payload ChangePasswordDto
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	authentication := c.Locals("authentication").(middlewares.Authentication)

	user, err := services.FindUserById(authentication.Id)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if !utils.CheckPasswordHash(payload.OldPassword, user.Password) {
		return c.Status(fiber.StatusForbidden).SendString("Wrong password")
	}

	user.Password, _ = utils.HashPassword(payload.NewPassword)
	database.GetConnection().Save(&user)

	return c.SendStatus(fiber.StatusOK)
}
