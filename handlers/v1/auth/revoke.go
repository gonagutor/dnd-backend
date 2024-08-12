package v1_auth_handlers

import (
	"dnd/backend/constants/http_codes"
	utils_constants "dnd/backend/constants/utils"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

//	@Tags					Auth
//  @Description	Revokes the refresh tokens
//	@Accept				json
//	@Produce			json
//  @Param				Authorization	header	string	true	"Access token with Bearer prefix"
//  @Success			201	{object}	responses.CorrectResponse	"If the response is successful you will receive simple code and message indicating that the token has been revoked"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed or the user has not verified their email yet"
//  @Failure			500	{object}	responses.FailureResponse	"If the new token secret could not be saved it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/revoke [post]
func Revoke(ctx *fiber.Ctx) error {
	user := protected.GetUserFromContext(ctx)

	user.RefreshKey = utils.GenerateRandomCode(utils_constants.REFRESH_KEY_LENGTH)
	tx := utils.PGConnection.Save(user)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    http_errors.COULD_NOT_REVOKE_REFRESH_TOKEN,
			"message": "Could not revoke token",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.REFRESH_TOKEN_REVOKED,
		"message": "Refresh token revoked",
	})
}
