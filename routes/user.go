package routes

import (
	user "dnd/backend/handlers/v1/user"
	"dnd/backend/middleware/protected"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	userRouter := router.Group("/user")

	userRouter.Get("/:user_id", protected.New(protected.Config{}), user.GetOne)
	userRouter.Get("/", protected.New(protected.Config{}), user.GetAll)
	userRouter.Put("/:user_id", protected.New(protected.Config{}), user.EditUser)
	userRouter.Delete("/:user_id", protected.New(protected.Config{}), user.DeleteUser)
}
