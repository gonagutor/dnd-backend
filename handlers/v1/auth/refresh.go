package v1_auth_handlers

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	auth_utils "dnd/backend/utils/auth"

	"github.com/gofiber/fiber/v2"
)

type RefreshResponse struct {
	AccessToken string `example:"" format:"jwt"`
}

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

//	@Tags					Auth
//  @Description	Uses the refresh token in the header to generate a new access token for the user
//	@Produce			json
//  @Param				Authorization	header	string	true	"Refresh token with Bearer prefix"
//  @Success			200	{object}	responses.CorrectResponse{data=RefreshResponse}	"If the response is successful you will receive the new accessToken inside the data field of the response"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse	"The API can answer with a 403 if the token has expired or is invalid"
//  @Failure			500	{object}	responses.FailureResponse	"If the new token could not be generated it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/refresh [post]
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
