package v1_character_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type HealBody struct {
	Heal int `json:"heal"`
}

func Heal(ctx *fiber.Ctx) error {
	heal := new(HealBody)
	err := ctx.BodyParser(heal)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Body could not be parsed",
		})
	}

	characterId := ctx.Params("character_id")
	character, err := models.FindCharacterByID(characterId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_NOT_FOUND,
			"message": "Character could not be found",
		})
	}

	previousHP := character.TemporaryHP
	character.TemporaryHP += heal.Heal
	saveError := utils.PGConnection.Save(character)
	if saveError.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_UPDATE_CHARACTER,
			"message": "Character could not heal",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.HEALED,
		"message": "Character HP healed",
		"data": fiber.Map{
			"previousHP": previousHP,
			"currentHP":  character.TemporaryHP,
		},
	})
}
