package v1_item_handler

import (
	"github.com/gofiber/fiber/v2"

	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"
)

// @Tags Item
// @Description Delete an item
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token with Bearer prefix"
// @Param item_id path string true "The id of the item you want to delete"
// @Success 200 {object} responses.CorrectResponse "If the response is successful you will receive a simple code and message indicating that the item has been deleted"
//
//	@Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//	@Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed. the user has not verified their email yet or (if the user is trying to delete public items) the user is not an admin"
//
// @Failure 404 {object} responses.FailureResponse "If the item is not found the API will answer with a 404 code"
// @Failure 500 {object} responses.FailureResponse "If the server fails to delete the item it will answer with a 500 code. Please report this error if you encounter it in production"
// @Router /v1/item/{item_id} [delete]
func Delete(ctx *fiber.Ctx) error {
	itemId := ctx.Params("item_id")
	itemUrl, err := models.FindItemByID(itemId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.ITEM_NOT_FOUND,
			"message": "Item id by '" + itemId + "' not found",
		})
	}

	userLocal := protected.GetUserFromContext(ctx)
	if userLocal.Role != "admin" && *itemUrl.User != userLocal.ID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You can only delete your own items",
		})
	}

	err = utils.PGConnection.Delete(&itemUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_DELETE_ITEM,
			"message": "Item could not be deleted",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.ITEM_DELETED,
		"message": "Item deleted",
	})
}
