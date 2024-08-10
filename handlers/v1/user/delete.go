package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

//	@Tags					User
//  @Description	Delete an user by id
//	@Accept				json
//	@Produce			json
//  @Param				Authorization	header	string	true	"Access token with Bearer prefix"
//	@Param				user_id				path		string	true	"User's id"
//  @Success			200	{object}	responses.CorrectResponse	"If the response is successful you will receive a simple code and message indicating that the user has been deleted"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed. the user has not verified their email yet or (if the user is trying to delete other than self) the user is not an admin"
//  @Failure			404	{object}	responses.FailureResponse	"If the user could not be found it will return a 404 code"
//  @Failure			500	{object}	responses.FailureResponse	"If the user could not be deleted it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/user/{user_id} [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")
	userUrl, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "User id by '" + userId + "' not found",
		})
	}

	userLocal := ctx.Locals("user").(*models.User)
	if userLocal.Role != "admin" && userLocal.ID != userUrl.ID {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You can only edit your own user",
		})
	}

	err = utils.PGConnection.Delete(&userUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_DELETE_USER,
			"message": "User could not be deleted",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USER_DELETED,
		"message": "User deleted",
	})
}
