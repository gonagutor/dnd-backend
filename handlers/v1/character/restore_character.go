package v1_character_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func RestoreCharacter(ctx *fiber.Ctx) error {
	user := protected.GetUserFromContext(ctx)
	characterId := ctx.Params("character_id")
	err := models.RestoreDeletedCharacter(characterId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_NOT_FOUND,
			"message": "Character could not be found",
		})
	}

	character, err := models.FindCharacterByID(characterId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_NOT_FOUND,
			"message": "Character could not be found",
		})
	}

	if user.ID != character.UserID {
		utils.PGConnection.Delete(character)
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_DELETE_NOT_ALLOWED,
			"message": "This is not your character",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.CHARACTER_RESTORED,
		"message": "Character restored",
	})
}
