package v1_auth_handlers

import (
	"dnd/backend/constants/http_codes"
	utils_constants "dnd/backend/constants/utils"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	auth_utils "dnd/backend/utils/auth"
	"dnd/backend/utils/validators"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RedeemRecoveryCodeRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

func validateAndParseRedeemRecoveryCodeParams(ctx *fiber.Ctx) (*RedeemRecoveryCodeRequest, error) {
	recoveryData := new(RedeemRecoveryCodeRequest)
	err := ctx.BodyParser(recoveryData)
	if err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Request could not be processed",
		})
	}
	validationError := validators.Validator.Struct(recoveryData)
	if validationError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": validationError.Error(),
		})
	}
	return recoveryData, nil
}

func redeemRecoveryCodePrechecks(ctx *fiber.Ctx, id string) (*models.User, error) {
	user, err := models.FindUserByID(id)
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
//  @Description	Uses the provided token to change the user's password
//	@Accept				json
//	@Produce			json
//  @Param				Body	body	RedeemRecoveryCodeRequest	true	"The received token and the new password"
//  @Success			200	{object}	responses.CorrectResponse	"If the response is successful you will receive simple code and message indicating that the passworc has been changed"
//  @Failure			400	{object}	responses.FailureResponse	"If a field is missing or the body couldn't be parsed the API will answer with a 400 code. In case a field is missing or has the incorrect format it will return the field which fails"
//  @Failure			403	{object}	responses.FailureResponse	"The API can answer with a 403 if the token has expired or is invalid"
//  @Failure			500	{object}	responses.FailureResponse	"If the hashed password could not be generated it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/recover-password [post]
func RedeemRecoveryCode(ctx *fiber.Ctx) error {
	recoveryData, validationResponseError := validateAndParseRedeemRecoveryCodeParams(ctx)
	if recoveryData == nil {
		return validationResponseError
	}

	id, validateTokenError := auth_utils.ValidateRecoverToken(recoveryData.Token)
	if validateTokenError != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   http_errors.BAD_RECOVER_TOKEN,
			"message": validateTokenError.Error(),
		})
	}

	user, prechecksResponseError := redeemRecoveryCodePrechecks(ctx, id)
	if user == nil {
		return prechecksResponseError
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), utils_constants.PASSWORD_COST)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_UPDATE_PASSWORD,
			"message": "Could not update password",
		})
	}
	user.Password = string(passwordHashed)
	utils.PGConnection.Save(user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.PASSWORD_CHANGED,
		"message": "Password successfully changed",
	})
}
