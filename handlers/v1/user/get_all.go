package v1_user_handler

import (
	"dnd/backend/models"
	"dnd/backend/utils"
	"math"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	if user.Role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "NOT_AN_ADMIN",
			"message": "You do not have permission to access this resource",
		})
	}

	users, err := models.GetAllUsers(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "COULD_NOT_GET_USERS",
			"message": "Users could not be retrieved",
		})
	}

	page, pageSize := utils.Pagination(ctx)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "USERS_FOUND",
		"message": "Users found",
		"pagination": fiber.Map{
			"page":       page,
			"maxPages":   math.Ceil(float64(models.CountUsers()) / float64(pageSize)),
			"pageSize":   pageSize,
			"totalPages": math.Ceil(float64(models.CountUsers()) / float64(pageSize)),
		},
		"data": users,
	})
}
