package routes

import (
	"github.com/gofiber/fiber/v2"

	"revosearch/backend/handlers/auth"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
}
