package routes

import (
	"github.com/gofiber/fiber/v2"

	"revosearch/backend/handlers/auth"
	"revosearch/backend/middleware/protected"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Post("/recover-password", auth.Refresh)
	app.Post("/refresh", auth.Refresh)
	app.Post("/revoke", protected.New(protected.Config{}), auth.Revoke)
	app.Post("/validate-email", auth.ValidateEmail)
}
