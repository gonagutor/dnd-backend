package v1_character_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils/validators"
	"math"

	"github.com/gofiber/fiber/v2"
)

type LevelUpBody struct {
	Strength     int `json:"strength" validate:"max=2, min=0"`
	Dexterity    int `json:"dexterity" validate:"max=2, min=0"`
	Constitution int `json:"constitution" validate:"max=2, min=0"`
	Intelligence int `json:"intelligence" validate:"max=2, min=0"`
	Wisdom       int `json:"wisdom" validate:"max=2, min=0"`
	Charisma     int `json:"charisma" validate:"max=2, min=0"`
}

func LevelUp(ctx *fiber.Ctx) error {
	levelUp := new(LevelUpBody)
	validationError := validators.Validator.Struct(levelUp)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": validationError.Error(),
		})
	}

	err := ctx.BodyParser(levelUp)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Body could not be parsed",
		})
	}

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

	xpTable := [19]int{300, 900, 2700, 6500, 14000, 23000, 34000, 48000, 64000, 85000, 100000, 120000, 140000, 165000, 195000, 225000, 265000, 305000, 355000}
	if character.XP < xpTable[character.Level] {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_CANT_LEVEL_UP,
			"message": "This character cant level up yet",
		})
	}

	if levelUp.Strength+levelUp.Dexterity+levelUp.Constitution+levelUp.Intelligence+levelUp.Wisdom+levelUp.Charisma > 2 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.CHARACTER_CANT_LEVEL_UP,
			"message": "Too many stats",
		})
	}

	newLevel := detectLevel(character.XP, character.Level, xpTable[character.Level:])
	if math.Floor(float64(newLevel/4)) > math.Floor(float64(character.Level/4)) {
		character.ProficiencyBonus += 1

		character.Strength += levelUp.Strength
		character.Dexterity += levelUp.Dexterity
		character.Constitution += levelUp.Constitution
		character.Intelligence += levelUp.Intelligence
		character.Wisdom += levelUp.Wisdom
		character.Charisma += levelUp.Charisma
	}

	character.Level = newLevel
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.LEVEL_UP,
		"message": "Character level up",
		"data": fiber.Map{
			"currentLevel": newLevel,
			"strength":     character.Strength,
			"dexterity":    character.Dexterity,
			"constitution": character.Constitution,
			"intelligence": character.Intelligence,
			"wisdom":       character.Wisdom,
			"charisma":     character.Charisma,
		},
	})
}

func detectLevel(xp int, currentLevel int, xpTable []int) int {
	newLevel := currentLevel
	for i := 0; i < len(xpTable); i++ {
		if xp >= xpTable[i] {
			newLevel++
		} else {
			break
		}
	}

	return newLevel
}
