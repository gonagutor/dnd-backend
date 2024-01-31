package v1_class_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	"dnd/backend/utils/validators"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func validateAndParseCreateClassBody(ctx *fiber.Ctx) (*models.Class, error) {
	ret := new(models.Class)
	bodyParserError := ctx.BodyParser(ret)
	if bodyParserError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Request could not be processed",
			"reason": bodyParserError.Error(),
		})
	}

	validationError := validators.Validator.Struct(ret)
	if validationError != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Validation failed for one or more fields",
			"reason": strings.Split(validationError.Error(), "\n"),
		})
	}

	return ret, nil
}

func CreateClass(ctx *fiber.Ctx) error {
	body, err := validateAndParseCreateClassBody(ctx)
	if body == nil {
		return err
	}

	_, err = utils.ClassCollection.InsertOne(ctx.Context(), body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_CREATE_CLASS,
			"message": "Class could not be created due to an internal error",
			"reason": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    http_codes.CLASS_CREATED,
		"message": "Class created successfully",
	})
}
