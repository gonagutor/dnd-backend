package routes

import (
	v1_character_handler "dnd/backend/handlers/v1/character"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupCharacterRoutes(router fiber.Router) {
	characters := router.Group("/character")

	characters.Patch("/:character_id/coins", v1_character_handler.UpdateCoins)
	characters.Delete("/:character_id/delete", protected.New(protected.Config{}), v1_character_handler.DeleteCharacter)
	characters.Put("/:character_id/restore_character", protected.New(protected.Config{}), v1_character_handler.RestoreCharacter)
}
