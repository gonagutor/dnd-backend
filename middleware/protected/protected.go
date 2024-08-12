package protected

import (
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	auth_utils "dnd/backend/utils/auth"

	"github.com/gofiber/fiber/v2"
)

type UserKey struct {
}

type Config struct {
	Filter func(ctx *fiber.Ctx) bool // Required
}

func GetUserFromContext(ctx *fiber.Ctx) *models.User {
	return ctx.Locals("user").(*models.User)
}

func New(config Config) fiber.Handler {
	cfg := &Config{
		Filter: config.Filter,
	}
	return func(ctx *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(ctx) {
			return ctx.Next()
		}

		token, tokenNotExtractable := utils.ExtractToken(ctx)
		if tokenNotExtractable != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   http_errors.BAD_REQUEST,
				"message": "Could not extract token",
			})
		}
		id, errorValidatingToken := auth_utils.ValidateAccessToken(token)
		if errorValidatingToken != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   http_errors.BAD_ACCESS_TOKEN,
				"message": "Access token has expired, was tampered with or is malformed",
			})
		}
		user, errorFindingUser := models.FindUserByID(id)
		if errorFindingUser != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   http_errors.BAD_ACCESS_TOKEN,
				"message": "User no longer exists",
			})
		}

		if !user.IsActive {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   http_errors.EMAIL_NOT_VERIFIED,
				"message": "Email not verified, please verify your email and try again",
			})
		}

		ctx.Locals("user", user)
		return ctx.Next()
	}
}
