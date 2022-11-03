package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Serbroda/ragbag/pkg/database"
	"github.com/Serbroda/ragbag/pkg/middlewares"
	"github.com/Serbroda/ragbag/pkg/models"
	"github.com/Serbroda/ragbag/pkg/services"
	"github.com/Serbroda/ragbag/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetMetaInfo(c *fiber.Ctx) error {
	link := c.Params("url")
	key := c.Query("key")
	meta, err := utils.Parse(link)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if key != "" {
		j, _ := json.Marshal(meta)
		value := gjson.Get(string(j), key)
		return c.Status(fiber.StatusOK).SendString(value.Str)
	}
	return c.Status(fiber.StatusOK).JSON(meta)
}

func GetLinks(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	group, err := services.FindGroupById(groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.Visibility == models.Private && group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	links := services.FindLinksByGroupId(groupId)

	return c.Status(fiber.StatusOK).JSON(links)
}

func GetLinksPublic(c *fiber.Ctx) error {
	groupId := c.Params("groupId")

	group, err := services.FindGroupById(groupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.Visibility != models.Public {
		return c.SendStatus(fiber.StatusForbidden)
	}

	links := services.FindLinksByGroupId(groupId)

	return c.Status(fiber.StatusOK).JSON(links)
}

func GetLatestLinks(c *fiber.Ctx) error {
	order := c.Query("order")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
	authentication := c.Locals("authentication").(middlewares.Authentication)

	links := services.FindLinks(authentication.Id, order, int(limit))

	return c.Status(fiber.StatusOK).JSON(links)
}

func GetLink(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func CreateLink(c *fiber.Ctx) error {
	groupId := c.Params("groupId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var dto models.CreateLinkDto
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

	link := models.Link{
		Name:        dto.Name,
		Url:         dto.Url,
		Description: dto.Description,
		GroupId:     groupId,
		Visibility:  models.Private,
	}

	site, err := utils.Parse(link.Url)

	if err == nil {
		link.Url = site.Url
		if link.Name == "" {
			link.Name = site.Meta.Title
		}
		if link.Name == "" {
			link.Name = site.Meta.SiteName
		}
		if link.Name == "" {
			link.Name = site.Host
		}
		if link.Description == "" {
			link.Description = strings.TrimSpace(site.Meta.Description)
		}
	}

	database.GetConnection().Create(&link)
	return c.Status(http.StatusCreated).JSON(link)
}

func UpdateLink(c *fiber.Ctx) error {
	linkId := c.Params("linkId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	var dto models.CreateLinkDto
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	link, err := services.FindLinkById(linkId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Link not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	group, err := services.FindGroupById(link.GroupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	if dto.GroupId != "" && dto.GroupId != link.GroupId {
		group, err := services.FindGroupById(dto.GroupId)

		if err != nil && err == services.ErrEntityNotFound {
			return c.Status(fiber.StatusNotFound).SendString("Group not found")
		} else if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		} else if group.OwnerId != authentication.Id {
			return c.SendStatus(fiber.StatusForbidden)
		}

		link.GroupId = dto.GroupId
	}

	if link.Url != dto.Url {
		link.Url = dto.Url
	}
	if link.Name != dto.Name {
		link.Name = dto.Name
	}
	if link.Description != dto.Description {
		link.Description = strings.TrimSpace(dto.Description)
	}

	database.GetConnection().Save(&link)

	return c.Status(fiber.StatusOK).JSON(link)
}

func DeleteLink(c *fiber.Ctx) error {
	linkId := c.Params("linkId")
	authentication := c.Locals("authentication").(middlewares.Authentication)

	link, err := services.FindLinkById(linkId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Link not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	group, err := services.FindGroupById(link.GroupId)

	if err != nil && err == services.ErrEntityNotFound {
		return c.Status(fiber.StatusNotFound).SendString("Group not found")
	} else if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if group.OwnerId != authentication.Id {
		return c.SendStatus(fiber.StatusForbidden)
	}

	database.GetConnection().Delete(&link)

	return c.SendStatus(fiber.StatusOK)
}
