package v1_character_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteCharacter(ctx *fiber.Ctx) error {
	user := protected.GetUserFromContext(ctx)
	characterId := ctx.Params("character_id")
	character, err := models.FindCharacterByID(characterId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_NOT_FOUND,
			"message": "Character could not be found",
		})
	}

	if user.ID != character.UserID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_DELETE_NOT_ALLOWED,
			"message": "This is not your character",
		})
	}

	saveError := utils.PGConnection.Delete(character)
	if saveError != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_DELETE_CHARACTER,
			"message": "Character could not be deleted",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.CHARACTER_DELETED,
		"message": "Character deleted",
		"data": fiber.Map{
			"id":      character.ID,
			"name":    character.Name,
			"classId": character.ClassID,
			"raceId":  character.RaceID,
		},
	})
}
