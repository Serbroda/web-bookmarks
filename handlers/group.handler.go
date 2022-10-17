package handlers

import (
	"net/http"
	"webcrate/database"
	"webcrate/middlewares"
	"webcrate/models"
	"webcrate/services"

	"github.com/gofiber/fiber/v2"
)

func GetGroups(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)
	return c.Status(fiber.StatusOK).JSON(services.FindGroupsByOwnerId(authentication.Id))
}

func GetGroup(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	group, err := services.FindGroupById(groupId)
	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(fiber.StatusOK).JSON(group)
}

func CreateGroup(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var dto models.CreateGroupDto
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var group models.Group
	result := database.GetConnection().Where("name = ?", dto.Name).Find(&group)

	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Group " + group.Name + " already exists")
	}

	group = models.Group{
		Icon:        dto.Icon,
		Name:        dto.Name,
		Description: dto.Description,
		OwnerId:     authentication.Id,
		Visibility:  models.Private,
	}
	database.GetConnection().Create(&group)

	return c.Status(http.StatusCreated).JSON(group)
}

func UpdateGroup(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var dto models.CreateGroupDto
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	group, err := services.FindGroupById(groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	if dto.Icon != "" {
		group.Icon = dto.Icon
	}
	if dto.Name != "" {
		group.Name = dto.Name
	}
	group.Description = dto.Description

	database.GetConnection().Save(&group)

	return c.Status(http.StatusCreated).JSON(group)
}

func DeleteGroup(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	group, err := services.FindGroupById(groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	database.GetConnection().Delete(&group)
	return c.SendStatus(fiber.StatusOK)
}
