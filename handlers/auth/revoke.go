package auth

import (
	"revosearch/backend/constants/http_codes"
	utils_constants "revosearch/backend/constants/utils"
	"revosearch/backend/errors/http_errors"
	"revosearch/backend/middleware/protected"
	"revosearch/backend/utils"

	"github.com/gofiber/fiber/v2"
)

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
