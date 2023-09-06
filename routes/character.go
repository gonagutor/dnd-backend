package routes

import (
	v1_character_handler "dnd/backend/handlers/v1/character"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupCharacterRoutes(app *fiber.App) {
	characters := app.Group("/character")

	characters.Patch("/:character_id/coins", v1_character_handler.UpdateCoins)
	characters.Delete("/:character_id/delete", protected.New(protected.Config{}))
}
