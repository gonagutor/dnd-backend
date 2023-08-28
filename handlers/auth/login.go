package auth

import (
	"github.com/gofiber/fiber/v2"

	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	auth_utils "dnd/backend/utils/auth"
	"dnd/backend/utils/validators"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

func validateAndParseLoginBody(ctx *fiber.Ctx) (*LoginRequest, error) {
	login := new(LoginRequest)
	err := ctx.BodyParser(login)
	if err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Request could not be processed",
		})
	}
	validationError := validators.Validator.Struct(login)
	if validationError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": validationError.Error(),
		})
	}
	return login, nil
}

func loginPrechecks(ctx *fiber.Ctx, body *LoginRequest) (*models.User, error) {
	user, err := models.FindUserByEmail(body.Email)
	if err != nil {
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL,
			"message": "Email could not be found. User not registered",
		})
	}
	passwordCorrect := user.CheckPassword(body.Password)
	if passwordCorrect != nil {
		println(passwordCorrect.Error())
		return nil, ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.BAD_PASSWORD,
			"message": "Incorrect password",
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

func Login(ctx *fiber.Ctx) error {
	login, validationResponseError := validateAndParseLoginBody(ctx)
	if login == nil {
		return validationResponseError
	}

	user, prechecksResponseError := loginPrechecks(ctx, login)
	if user == nil {
		return prechecksResponseError
	}

	accessToken, accessTokenError := auth_utils.GenerateAccessToken(user)
	refreshToken, refreshTokenError := auth_utils.GenerateRefreshToken(user)
	if accessTokenError != nil || refreshTokenError != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_GENERATE_TOKEN,
			"message": "Token could not be generated",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.LOGGED_IN,
		"message": "Logged in correctly",
		"data": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}
