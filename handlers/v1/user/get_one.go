package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"

	"github.com/gofiber/fiber/v2"
)

//	@Tags					User
//  @Description	Retrieve an user by id
//	@Accept				json
//	@Produce			json
//  @Param				Authorization	header	string	true	"Access token with Bearer prefix"
//	@Param				user_id				path		string	true	"User's id"
//  @Success			200	{object}	responses.CorrectResponse{data=models.User}	"If the response is successful you will receive the requested user inside the data key"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed or the user has not verified their email yet"
//  @Failure			404	{object}	responses.FailureResponse	"If the if the user could not be found it will return a 404 code"
//  @Router 		/v1/user/{user_id} [get]
func GetOne(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")

	user, err := models.FindUserByID(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   http_errors.USER_NOT_FOUND,
			"message": "User with id '" + userId + "' not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USER_FOUND,
		"message": "User found",
		"data": fiber.Map{
			"id":         user.ID,
			"email":          user.Email,
			"name":           user.Name,
			"surname":        user.Surname,
			"role":           user.Role,
			"profilePicture": user.ProfilePicture,
			"isActive":       user.IsActive,

			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}
