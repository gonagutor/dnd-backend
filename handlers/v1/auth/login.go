package v1_auth_handlers

import (
	"github.com/gofiber/fiber/v2"

	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	auth_utils "dnd/backend/utils/auth"
	"dnd/backend/utils/validators"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john@doe.com" format:"email"`
	Password string `json:"password" validate:"required,min=8,max=64" example:"Testtest1@"`
}

type NonSensitiveDataUser struct {
	Id string `json:"id" example:"eb7ed8ed-3316-47f3-a8c5-a0c17aa147f4"`
	Name string `json:"name" example:"Gonzalo"`
	Surname string `json:"surname" example:"Aguado Torres"`
	ProfilePicture string `json:"profilePicture" example:"https://picsum.photos/200/300"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzIiwiaXNzIjoiZG5kIiwic3ViIjoiMDczZDhmN2ItMDIyNS00NGRjLWE3NTMtOGJmNzYyYmVkMzc0IiwiZXhwIjoxNzA2NDQwNjExfQ.YQ6shB0HGGw9tN5jo6cBzqjoB4LxGlNadC52exF_Hm7UYfsbf8uB-u1Sq7ukgkIkkHw-eR0VLwmjNCWmWoF6tA"`
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJrZXkiOiIkMmEkMTAkQmYwZ2FlUmxkOXRzWVEzbnBXSkFBZVlhdFlFZG9yWVVtRFl5VHl6TVA0a0lSRHdna1B3Y2EiLCJ0eXBlIjoicmVmcmVzaCIsImlzcyI6ImRuZCIsInN1YiI6IjA3M2Q4ZjdiLTAyMjUtNDRkYy1hNzUzLThiZjc2MmJlZDM3NCJ9.pv76u4p-kfAwGu8VPwzAKv5lGclrI85T2Uuu0kCT24hlfLRnjpU7iktgtlPujWuB_NVHxBKlvz_qkmyWeqLxlw"`
	User NonSensitiveDataUser `json:"user"`
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

//	@Tags					Auth
//  @Description	Login request that returns a refresh and an access token
//	@Accept				json
//	@Produce			json
//  @Param				Body	body	LoginRequest	true	"Simple password and email login"
//  @Success			200	{object}	responses.CorrectResponse{data=LoginResponse}	"When the response is successful you will receive an access token and a refresh token, along with some basic user data"
//  @Failure			400	{object}	responses.FailureResponse											"If a field is missing or the body couldn't be parsed the API will answer with a 400 code. In case a field is missing or has the incorrect format it will return the field which fails"
//  @Failure			403	{object}	responses.FailureResponse											"The API can answer with a 403 if the password is incorrect or if the email has not been verified yet"
//  @Failure			500	{object}	responses.FailureResponse											"If either token could not be generated it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/login [post]
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
			"user": fiber.Map{
				"id":             user.ID,
				"name":           user.Name,
				"surname":        user.Surname,
				"profilePicture": user.ProfilePicture,
			},
		},
	})
}
