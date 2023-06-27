package auth

import (
	"revosearch/backend/constants/http_codes"
	"revosearch/backend/errors/http_errors"
	"revosearch/backend/models"
	"revosearch/backend/utils"
	auth_utils "revosearch/backend/utils/auth"

	"github.com/gofiber/fiber/v2"
)

type ValidateEmailQuery struct {
	Token string `json:"token" validate:"required"`
}

func ValidateEmail(ctx *fiber.Ctx) error {
	emailQuery := new(ValidateEmailQuery)
	badQuery := ctx.QueryParser(emailQuery)
	if badQuery != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL_TOKEN,
			"message": "No email token was provided",
		})
	}

	id, errorValidatingToken := auth_utils.ValidateEmailToken(emailQuery.Token)
	if errorValidatingToken != nil {
		print(errorValidatingToken.Error())
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_ACCESS_TOKEN,
			"message": "Provided email token was malformed or has expired",
		})
	}
	user, err := models.FindUserByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "An error happend when finding the user or the user for which this token was generated was deleted",
		})
	}
	if user.IsActive {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   http_errors.EMAIL_ALREADY_VERIFIED,
			"message": "Your email has already been verified",
		})
	}

	user.IsActive = true
	tx := utils.PGConnection.Save(user)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_VERIFY_EMAIL,
			"message": "Could not verify your email",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.EMAIL_VERIFIED,
		"message": "Email was verified correctly",
	})
}
