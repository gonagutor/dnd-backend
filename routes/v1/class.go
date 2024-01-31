package v1_routes

import (
	v1_class_handler "dnd/backend/handlers/v1/class"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupClassRoutes(v1 fiber.Router) {
	class := v1.Group("/class", protected.New(protected.Config{}))

	class.Post("/", v1_class_handler.CreateClass)
	class.Get("/", v1_class_handler.ClassIndex)
	class.Get("/:id", v1_class_handler.GetClassById)
	class.Put("/:id", v1_class_handler.EditClass)
}
