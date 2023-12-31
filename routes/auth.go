package routes

import (
	"github.com/gofiber/fiber/v2"

	"dnd/backend/handlers/auth"
	"dnd/backend/middleware/protected"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Post("/refresh", auth.Refresh)
	app.Post("/recover-password", auth.RedeemRecoveryCode)
	app.Post("/recover-password-request", auth.RecoverPassword)
	app.Post("/revoke", protected.New(protected.Config{}), auth.Revoke)
	app.Get("/validate-email", auth.ValidateEmail)
}
