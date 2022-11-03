package handlers

import (
	"net/http"
	"strconv"

	"github.com/Serbroda/ragbag/database"
	"github.com/Serbroda/ragbag/middlewares"
	"github.com/Serbroda/ragbag/models"
	"github.com/Serbroda/ragbag/services"

	"github.com/gofiber/fiber/v2"
)

func GetGroups(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)
	return c.Status(fiber.StatusOK).JSON(services.FindGroupsByOwnerId(authentication.Id))
}

func GetLatestGroups(c *fiber.Ctx) error {
	order := c.Query("order")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	authentication := c.Locals("authentication").(middlewares.Authentication)
	return c.Status(fiber.StatusOK).JSON(services.FindLatestGroups(authentication.Id, order, int(limit)))
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

	if group.Visibility != models.Public && group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(fiber.StatusOK).JSON(group)
}

func GetGroupPublic(c *fiber.Ctx) error {
	groupId := c.Params("groupId")

	group, err := services.FindGroupById(groupId)
	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if group.Visibility != models.Public {
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
		return c.Status(fiber.StatusConflict).SendString("Group " + group.Name + " already exists")
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

func ChangeGroupVisibility(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var dto models.ChangeGroupVisibility
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	group, err := services.FindGroupById(groupId)
	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	group.Visibility = dto.Visibility
	database.GetConnection().Save(&group)

	return c.Status(fiber.StatusOK).JSON(group)
}

func GetGroupSubscriptions(c *fiber.Ctx) error {
	authentication := c.Locals("authentication").(middlewares.Authentication)
	return c.Status(fiber.StatusOK).JSON(services.FindGroupSubscriptions(authentication.Id))
}

func CreateGroupSubscription(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	group, err := services.FindGroupById(groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.Visibility != models.Public {
		return c.SendStatus(fiber.StatusForbidden)
	}

	subscription, err := services.FindGroupSubscription(authentication.Id, groupId)
	if err == nil {
		return c.Status(fiber.StatusConflict).SendString("Subscription already exists")
	}

	subscription = models.GroupSubscription{
		GroupId: groupId,
		UserId:  authentication.Id,
	}

	database.GetConnection().Create(&subscription)

	return c.Status(http.StatusCreated).JSON(subscription)
}

func DeleteGroupSubscription(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	subscription, err := services.FindGroupSubscription(authentication.Id, groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group subscription not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if subscription.UserId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	database.GetConnection().Delete(&subscription)
	return c.SendStatus(fiber.StatusOK)
}
