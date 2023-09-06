package routes

import (
	v1_character_handler "dnd/backend/handlers/v1/character"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupCharacterRoutes(app *fiber.App) {
	characters := app.Group("/character", protected.New(protected.Config{}))

	characters.Patch("/:character_id/coins", v1_character_handler.UpdateCoins)
	characters.Delete("/:character_id/delete", v1_character_handler.DeleteCharacter)
	characters.Put("/:character_id/restore_character", v1_character_handler.RestoreCharacter)
	characters.Patch("/:character_id/level_up", v1_character_handler.LevelUp)
}
