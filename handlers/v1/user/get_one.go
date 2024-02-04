package v1_user_handler

import (
	"dnd/backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetOne(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")

	user, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "USER_NOT_FOUND",
			"message": "User id by '" + userId + "' not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "USER_FOUND",
		"message": "User found",
		"data": fiber.Map{
			"userId":         user.ID,
			"email":          user.Email,
			"name":           user.Name,
			"surname":        user.Surname,
			"role":           user.Role,
			"profilePicture": user.ProfilePicture,
			"isActive":       user.IsActive,

			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}
