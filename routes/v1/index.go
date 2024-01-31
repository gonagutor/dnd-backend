package v1_routes

import (
	"dnd/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupVersionedRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/", handlers.V1Index)

	SetupClassRoutes(v1)
}
