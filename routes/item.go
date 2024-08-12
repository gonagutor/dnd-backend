package routes

import (
	"github.com/gofiber/fiber/v2"

	item "dnd/backend/handlers/v1/item"
	"dnd/backend/middleware/protected"
)

func SetupItemRoutes(router fiber.Router) {
	itemRouter := router.Group("/item")

	itemRouter.Post("/", protected.New(protected.Config{}), item.Create)
	itemRouter.Put("/:item_id", protected.New(protected.Config{}), item.Edit)
	itemRouter.Delete("/:item_id", protected.New(protected.Config{}), item.Delete)
}