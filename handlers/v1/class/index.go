package v1_class_handler

import "github.com/gofiber/fiber/v2"

func ClassIndex(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"get_class_by_id": "/v1/class/:id",
		"edit_class_by_id": "/v1/class/:id",
		"delete_class": "/v1/class",
		"create_class": "/v1/class",
	})
}
