package v1_auth_handlers

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	auth_utils "dnd/backend/utils/auth"
	"dnd/backend/utils/validators"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" format:"email" example:"john@doe.com"`
	Password string `json:"password" validate:"required,password" example:"Testtest1@"`
	Name     string `json:"name" validate:"required,max=32" example:"Gonzalo"`
	Surname  string `json:"surname" validate:"required,max=64" example:"Aguado Torres"`
}

func validateAndParseRegisterBody(ctx *fiber.Ctx) (*RegisterRequest, error) {
	ret := new(RegisterRequest)
	bodyParserError := ctx.BodyParser(ret)
	if bodyParserError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Request could not be processed",
		})
	}

	validationError := validators.Validator.Struct(ret)
	if validationError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": validationError.Error(),
		})
	}

	return ret, nil
}

func registrationPrechecks(ctx *fiber.Ctx, body *RegisterRequest) (bool, error) {
	userExists, _ := models.FindUserByEmail(body.Email)
	if userExists != nil {
		return false, ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   http_errors.BAD_EMAIL,
			"message": "Email already in use",
		})
	}

	return true, nil
}

func createUser(ctx *fiber.Ctx, body *RegisterRequest) (*models.User, error) {
	user := &models.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
		Surname:  body.Surname,
	}
	creationError := utils.PGConnection.Create(user).Error
	if creationError != nil {
		return nil, ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_CREATE_USER,
			"message": "Could not create user",
		})
	}

	return user, nil
}

//	@Tags					Auth
//  @Description	Creates an account and sends an email for verification
//	@Accept				json
//	@Produce			json
//  @Param				Body	body	RegisterRequest	true	"Email, name, surname and password required for registration"
//  @Success			201	{object}	responses.CorrectResponse	"If the response is successful you will receive simple code and message indicating that the account has been created"
//  @Failure			400	{object}	responses.FailureResponse	"If a field is missing or the body couldn't be parsed the API will answer with a 400 code. In case a field is missing or has the incorrect format it will return the field which fails"
//  @Failure			409 {object}	responses.FailureResponse "If the email already exists the server will return a 409 code"
//  @Failure			500	{object}	responses.FailureResponse	"If the verification token could not be generated or the user could not be created it will return a 500 code. Please report this error if you encounter it in production"
//  @Failure			502	{object}	responses.FailureResponse	"If nothing failed but the email could not be sent the server will return a 502 code. Please report this error if you encounter it in production"
//  @Router 		/v1/auth/register [post]
func Register(ctx *fiber.Ctx) error {
	body, err := validateAndParseRegisterBody(ctx)
	if body == nil {
		return err
	}

	passesPrechecks, prechecksResponseError := registrationPrechecks(ctx, body)
	if !passesPrechecks {
		return prechecksResponseError
	}

	user, userResponseError := createUser(ctx, body)
	if user == nil {
		return userResponseError
	}

	token, tokenErr := auth_utils.GenerateEmailToken(user.ID.String())
	if tokenErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_GENERATE_VERIFICATION_CODE,
			"message": "Verification code could not be generated",
		})
	}

	emailError := utils.SendMail("validate_email", user.Email, "Verifica tu cuenta en dnd", fiber.Map{
		"name":  user.Name,
		"token": token,
	})
	if emailError != nil {
		log.Println(emailError)
		utils.PGConnection.Delete(user)
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_SEND_EMAIL,
			"message": "Email could not be sent",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    http_codes.USER_CREATED,
		"message": "User created successfully",
	})
}
