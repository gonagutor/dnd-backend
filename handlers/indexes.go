package handlers

import "github.com/gofiber/fiber/v2"

var v1Routes = fiber.Map{
	"index": "/v1/",
}

var routes = fiber.Map{
	"login":            "/login",
	"register":         "/register",
	"refresh":          "/refresh",
	"revoke":           "/revoke",
	"recover_password": "/recover-password",
	"validate_email":   "/validate-email",
	"v1":               v1Routes,
}

func APIIndex(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"index": "/v1",
	})
}

func V1Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{})
}

func Status(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "running",
		"routes": routes,
	})
}
