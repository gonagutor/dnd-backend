package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")
	userUrl, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "User id by '" + userId + "' not found",
		})
	}

	userLocal := ctx.Locals("user").(*models.User)
	if userLocal.Role != "admin" && userLocal.ID != userUrl.ID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You can only edit your own user",
		})
	}

	err = utils.PGConnection.Delete(&userUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_EDIT_USER,
			"message": "User could not be edited",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USER_DELETED,
		"message": "User deleted",
	})
}
