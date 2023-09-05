package v1_character_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateCoinsBody struct {
	Copper   int `json:"copper"`
	Silver   int `json:"silver"`
	Electrum int `json:"electrum"`
	Gold     int `json:"gold"`
	Platinum int `json:"platinum"`
}

func UpdateCoins(ctx *fiber.Ctx) error {
	coins := new(UpdateCoinsBody)
	err := ctx.BodyParser(coins)
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

	previousCopper := character.CopperCoins
	previousSilver := character.SilverCoins
	previousElectrum := character.ElectrumCoins
	previousGold := character.GoldCoins
	previousPlatinum := character.PlatinumCoins
	character.CopperCoins += coins.Copper
	character.SilverCoins += coins.Silver
	character.ElectrumCoins += coins.Electrum
	character.GoldCoins += coins.Gold
	character.PlatinumCoins += coins.Platinum
	saveError := utils.PGConnection.Save(character)
	if saveError.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_UPDATE_CHARACTER,
			"message": "Character could not be updated",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.COINS_UPDATED,
		"message": "Character coins updated",
		"data": fiber.Map{
			"previousCopper":   previousCopper,
			"copper":           character.CopperCoins,
			"previousSilver":   previousSilver,
			"silver":           character.SilverCoins,
			"previousElectrum": previousElectrum,
			"electrum":         character.ElectrumCoins,
			"previousGold":     previousGold,
			"gold":             character.GoldCoins,
			"previousPlatinum": previousPlatinum,
			"platinum":         character.PlatinumCoins,
		},
	})
}
