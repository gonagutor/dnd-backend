package auth

import "github.com/gofiber/fiber/v2"

func Refresh(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{})
}
