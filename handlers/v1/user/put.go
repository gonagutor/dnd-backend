package v1_user_handler

import (
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserBody struct {
	Name           *string
	Surname        *string
	Role           *string
	ProfilePicture *string
	IsActive       *bool
}

func EditUser(ctx *fiber.Ctx) error {
	editUser := new(UpdateUserBody)
	err := ctx.BodyParser(editUser)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Body could not be parsed",
		})
	}

	userId := ctx.Params("user_id")
	userUrl, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "USER_NOT_FOUND",
			"message": "User id by '" + userId + "' not found",
		})
	}

	userLocal := ctx.Locals("user").(*models.User)
	if userLocal.Role != "admin" && userLocal.ID != userUrl.ID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "NOT_AN_ADMIN",
			"message": "You can only edit your own user",
		})
	}

	if userLocal.Role != "admin" && editUser.Role != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "NOT_AN_ADMIN",
			"message": "You can not edit the role",
		})
	}

	if editUser.Name != nil {
		userUrl.Name = *editUser.Name
	}
	if editUser.Surname != nil {
		userUrl.Surname = *editUser.Surname
	}
	if editUser.Role != nil {
		userUrl.Role = *editUser.Role
	}
	if editUser.ProfilePicture != nil {
		userUrl.ProfilePicture = *editUser.ProfilePicture
	}
	if editUser.IsActive != nil {
		userUrl.IsActive = *editUser.IsActive
	}

	err = utils.PGConnection.Save(&userUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "COULD_NOT_EDIT_USER",
			"message": "User could not be edited",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "USER_EDITED",
		"message": "User edited",
	})
}
