package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetOne(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")

	user, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "User with id '" + userId + "' not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USER_FOUND,
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
