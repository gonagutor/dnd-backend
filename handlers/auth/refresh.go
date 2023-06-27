package auth

import (
	"revosearch/backend/constants/http_codes"
	"revosearch/backend/errors/http_errors"
	"revosearch/backend/models"
	"revosearch/backend/utils"
	auth_utils "revosearch/backend/utils/auth"

	"github.com/gofiber/fiber/v2"
)

func refreshPrechecks(ctx *fiber.Ctx, refreshToken string) (*models.User, error) {
	userId, key, errorValidatingToken := auth_utils.ValidateRefreshToken(refreshToken)
	if errorValidatingToken != nil {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_REFRESH_TOKEN,
			"message": "Refresh token was tampered with or is malformed",
		})
	}
	user, errorFindingUser := models.FindUserByID(userId)
	if errorFindingUser != nil {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_REFRESH_TOKEN,
			"message": "User no longer exists",
		})
	}
	keyError := user.CheckKey(key)
	if keyError != nil {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.REVOKED_REFRESH_TOKEN,
			"message": "Refresh token has been revoked",
		})
	}
	return user, nil
}

func Refresh(ctx *fiber.Ctx) error {
	refreshToken, err := utils.ExtractToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Malformed token or Authorization header",
		})
	}

	user, precheckResponseError := refreshPrechecks(ctx, refreshToken)
	if user == nil {
		return precheckResponseError
	}

	token, tokenError := auth_utils.GenerateAccessToken(user)
	if tokenError != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_GENERATE_TOKEN,
			"message": "Token could not be generated",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.TOKEN_REFRESHED,
		"message": "Token refreshed correctly",
		"data": fiber.Map{
			"accessToken": token,
		},
	})
}
