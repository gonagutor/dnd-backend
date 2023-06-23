package handlers

import "github.com/gofiber/fiber/v2"

func APIIndex(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"v1": "/api/v1",
	})
}

func V1Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{})
}

func Status(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "running",
	})
}
