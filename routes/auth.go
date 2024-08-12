package routes

import (
	"github.com/gofiber/fiber/v2"

	auth "dnd/backend/handlers/v1/auth"
	"dnd/backend/middleware/protected"
)

func SetupAuthRoutes(router fiber.Router) {
	authGroup := router.Group("auth");

	authGroup.Post("/login", auth.Login)
	authGroup.Post("/register", auth.Register)
	authGroup.Post("/refresh", auth.Refresh)
	authGroup.Post("/recover-password", auth.RedeemRecoveryCode)
	authGroup.Post("/recover-password-request", auth.RecoverPassword)
	authGroup.Post("/revoke", protected.New(protected.Config{}), auth.Revoke)
	authGroup.Get("/validate-email", auth.ValidateEmail)
}
