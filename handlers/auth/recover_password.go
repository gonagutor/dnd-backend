package auth

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	auth_utils "dnd/backend/utils/auth"
	"dnd/backend/utils/validators"

	"github.com/gofiber/fiber/v2"
)

type RecoverPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func validateAndParseRecoverPasswordBody(ctx *fiber.Ctx) (*RecoverPasswordRequest, error) {
	recoverPassword := new(RecoverPasswordRequest)
	err := ctx.BodyParser(recoverPassword)
	if err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Request could not be processed",
		})
	}
	validationError := validators.Validator.Struct(recoverPassword)
	if validationError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": validationError.Error(),
		})
	}
	return recoverPassword, nil
}

func recoverPasswordPrechecks(ctx *fiber.Ctx, body *RecoverPasswordRequest) (*models.User, error) {
	user, err := models.FindUserByEmail(body.Email)
	if err != nil {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL,
			"message": "Email could not be found. User not registered",
		})
	}
	if !user.IsActive {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.EMAIL_NOT_VERIFIED,
			"message": "Email not verified, please verify your email and try again",
		})
	}
	return user, nil
}

func RecoverPassword(ctx *fiber.Ctx) error {
	recoverPassword, validationResponseError := validateAndParseRecoverPasswordBody(ctx)
	if recoverPassword == nil {
		return validationResponseError
	}

	user, prechecksResponseError := recoverPasswordPrechecks(ctx, recoverPassword)
	if user == nil {
		return prechecksResponseError
	}

	recoverToken, err := auth_utils.GenerateRecoverToken(user.ID.String())
	if err == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   http_errors.BAD_RECOVER_TOKEN,
			"message": err.Error(),
		})
	}

	emailError := utils.SendMail("recover_password", user.Email, "Recupera tu contraseña", fiber.Map{
		"name":  user.Name,
		"token": recoverToken,
	})
	if emailError != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_SEND_EMAIL,
			"message": "Email could not be sent",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.RECOVERY_EMAIL_SENT,
		"message": "Recovery email sent",
	})
}
