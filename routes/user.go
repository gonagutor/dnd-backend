package routes

import (
	v1_user_handler "dnd/backend/handlers/v1/user"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/pages", protected.New(protected.Config{}), v1_user_handler.GetPages)
	user.Get("/:user_id", protected.New(protected.Config{}), v1_user_handler.GetOne)
	user.Get("/", protected.New(protected.Config{}), v1_user_handler.GetAll)
	user.Put("/:user_id", protected.New(protected.Config{}), v1_user_handler.EditUser)
	user.Delete("/:user_id", protected.New(protected.Config{}), v1_user_handler.DeleteUser)
}
