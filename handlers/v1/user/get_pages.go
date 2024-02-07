package v1_user_handler

import (
	"dnd/backend/models"
	"dnd/backend/utils"
	"math"

	"github.com/gofiber/fiber/v2"
)

func GetPages(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	if user.Role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "NOT_AN_ADMIN",
			"message": "You do not have permission to access this resource",
		})
	}

	page, pageSize := utils.Pagination(ctx)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "PAGES_FOUND",
		"message": "Pages found",
		"pagination": fiber.Map{
			"page":     page,
			"maxPages": math.Ceil(float64(models.CountUsers()) / float64(pageSize)),
			"pageSize": pageSize,
		},
	})
}
