package v1_item_handler

import (
	"math"

	"github.com/gofiber/fiber/v2"

	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"
)

// @Tags Item
// @Description Get all items
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token with Bearer prefix"
// @Param page query int false "The page you want to retrieve"
// @Param pageSize query int false "The size of the page you want to retrieve"
// @Success 200 {object} responses.CorrectResponse{data=[]models.Item, pagination=responses.Pagination} "If the response is successful you will receive the items and pagination data"
// @Failure 400 {object} responses.FailureResponse "If no token is provided the API will answer with a 400 code"
// @Failure 403 {object} responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed, the user has not verified their email yet or its not an admin"
// @Failure 500 {object} responses.FailureResponse "If the server fails to get the items it will answer with a 500 code. Please report this error if you encounter it in production"
// @Router /v1/item/ [get]
func GetAll(ctx *fiber.Ctx) error {
	userLocal := protected.GetUserFromContext(ctx)
	if userLocal.Role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You do not have permission to access this resource",
		})
	}

	items, err := models.GetAllItems(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_GET_ITEMS,
			"message": "Items could not be retrieved",
		})
	}

	page, pageSize := utils.Pagination(ctx)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.ITEMS_FOUND,
		"message": "Items found",
		"pagination": fiber.Map{
			"page":     page,
			"maxPages": math.Ceil(float64(models.CountItems()) / float64(pageSize)),
			"pageSize": pageSize,
		},
		"data": items,
	})
}
