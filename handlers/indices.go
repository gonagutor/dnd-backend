package handlers

import "github.com/gofiber/fiber/v2"

var authRoutes = fiber.Map{
	"login":                    "/v1/login",
	"register":                 "/v1/register",
	"refresh":                  "/v1/refresh",
	"revoke":                   "/v1/revoke",
	"recover-password-request": "/v1/recover-password-request",
	"recover-password":         "/v1/recover-password",
	"validate-email":           "/v1/validate-email",
}

var v1Routes = fiber.Map{
	"index": "/v1/",
	"auth": authRoutes,
}

var routes = fiber.Map{
	"v1": v1Routes,
}

func APIIndex(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(routes)
}

func V1Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(v1Routes)
}

func Status(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "running",
		"routes": routes,
	})
}
