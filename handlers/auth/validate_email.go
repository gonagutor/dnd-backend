package auth

import (
	"revosearch/backend/constants/http_codes"
	"revosearch/backend/errors/http_errors"
	"revosearch/backend/models"
	"revosearch/backend/utils"
	auth_utils "revosearch/backend/utils/auth"
	"revosearch/backend/utils/validators"

	"github.com/gofiber/fiber/v2"
)

type ValidateEmailQuery struct {
	Token string `json:"token" validate:"required"`
}

func validateAndParseEmailQuery(ctx *fiber.Ctx) (*ValidateEmailQuery, error) {
	emailQuery := new(ValidateEmailQuery)
	badQuery := ctx.QueryParser(emailQuery)
	if badQuery != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL_TOKEN,
			"message": "No email token was provided",
		})
	}

	err := validators.Validator.Struct(emailQuery)
	if err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL_TOKEN,
			"message": "Token was malformed",
		})
	}
	return emailQuery, nil
}

func validatePrechecks(ctx *fiber.Ctx, token string) (*models.User, error) {
	id, errorValidatingToken := auth_utils.ValidateEmailToken(token)
	if errorValidatingToken != nil {
		print(errorValidatingToken.Error())
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_ACCESS_TOKEN,
			"message": "Provided email token was malformed or has expired",
		})
	}
	user, err := models.FindUserByID(id)
	if err != nil {
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "An error happend when finding the user or the user for which this token was generated was deleted",
		})
	}
	if user.IsActive {
		return nil, ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   http_errors.EMAIL_ALREADY_VERIFIED,
			"message": "Your email has already been verified",
		})
	}
	return user, nil
}

func ValidateEmail(ctx *fiber.Ctx) error {
	emailQuery, parseResponseError := validateAndParseEmailQuery(ctx)
	if emailQuery == nil {
		return parseResponseError
	}

	user, prechecksResponseError := validatePrechecks(ctx, emailQuery.Token)
	if user == nil {
		return prechecksResponseError
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
