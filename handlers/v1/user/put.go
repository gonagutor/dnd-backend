package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserBody struct {
	Name           *string	`json:"name" example:"Gonzalo"`
	Surname        *string	`json:"surname" example:"Aguado Torres"`
	Role           *string	`json:"role" example:"user"`
	ProfilePicture *string	`json:"profilePicture" example:"https://picsum.photos/200/300"`
	IsActive       *bool		`json:"isActive" example:"false"`
}

//	@Tags					User
//  @Description	Update an user by id
//	@Accept				json
//	@Produce			json
//  @Param				Authorization	header	string	true	"Access token with Bearer prefix"
//	@Param				user_id				path		string	true	"User's id"
//  @Param				Body					body		UpdateUserBody true "Fields to edit. Every field is optional. If the user is not an admin isActive and role are ignored"
//  @Success			200	{object}	responses.CorrectResponse	"If the response is successful you will receive a simple code and message indicating that the user has been edited"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed. the user has not verified their email yet or (if the user is trying to edit other than self) the user is not an admin"
//  @Failure			404	{object}	responses.FailureResponse	"If the user could not be found it will return a 404 code"
//  @Failure			500	{object}	responses.FailureResponse	"If the user could not be edited it will return a 500 code. Please report this error if you encounter it in production"
//  @Router 		/v1/user/{user_id} [put]
func EditUser(ctx *fiber.Ctx) error {
	editUser := new(UpdateUserBody)
	err := ctx.BodyParser(editUser)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Body could not be parsed",
		})
	}

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

	if userLocal.Role != "admin" && editUser.Role != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You can not edit the role",
		})
	}

	if editUser.Name != nil {
		userUrl.Name = *editUser.Name
	}
	if editUser.Surname != nil {
		userUrl.Surname = *editUser.Surname
	}
	if editUser.Role != nil && userLocal.Role == "admin" {
		userUrl.Role = *editUser.Role
	}
	if editUser.ProfilePicture != nil {
		userUrl.ProfilePicture = *editUser.ProfilePicture
	}
	if editUser.IsActive != nil && userLocal.Role == "admin" {
		userUrl.IsActive = *editUser.IsActive
	}

	err = utils.PGConnection.Save(&userUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_EDIT_USER,
			"message": "User could not be edited",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USER_EDITED,
		"message": "User edited",
	})
}
