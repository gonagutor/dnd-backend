package routes

import (
	item "dnd/backend/handlers/v1/item"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(router fiber.Router) {
	itemRouter := router.Group("/item")

	itemRouter.Post("/", protected.New(protected.Config{}), item.Create)
}
