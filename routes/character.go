package routes

import (
	character "dnd/backend/handlers/v1/character"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupCharacterRoutes(router fiber.Router) {
	charactersRouter := router.Group("/character")

	charactersRouter.Patch("/:character_id/coins", character.UpdateCoins)
	charactersRouter.Delete("/:character_id/delete", protected.New(protected.Config{}), character.DeleteCharacter)
	charactersRouter.Put("/:character_id/restore_character", protected.New(protected.Config{}), character.RestoreCharacter)
}
