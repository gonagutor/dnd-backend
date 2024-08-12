package v1_auth_handlers

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
	Email string `json:"email" validate:"required,email" example:"john@doe.com" format:"email"`
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

//	@Tags					Auth
//  @Description	Uses the provided email to send a password recovery email
//	@Accept				json
//	@Produce			json
//  @Param				Body	body	RecoverPasswordRequest	true	"Email inside json body"
//  @Success			200	{object}	responses.CorrectResponse	"If the response is successful you will receive simple code and message indicating that the email has been sent"
//  @Failure			400	{object}	responses.FailureResponse	"If a field is missing or the body couldn't be parsed the API will answer with a 400 code. In case a field is missing or has the incorrect format it will return the field which fails"
//  @Failure			403	{object}	responses.FailureResponse	"The API can answer with a 403 if the email doesn't exist or if the user has not verified their email"
//  @Failure			500	{object}	responses.FailureResponse	"If the token could not be generated it will return a 500 code. Please report this error if you encounter it in production"
//  @Failure			502	{object}	responses.FailureResponse	"If nothing failed but the email could not be sent the server will return a 502 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/recover-password-request [post]
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
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
