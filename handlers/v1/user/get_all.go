package v1_user_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/models"
	"dnd/backend/utils"
	"math"

	"github.com/gofiber/fiber/v2"
)

//	@Tags					User
//  @Description	Retrieve an user by id
//	@Accept				json
//	@Produce			json
//  @Param				Authorization	header	string	true	"Access token with Bearer prefix"
//	@Param				page					query		int			false "Page to show"
//  @Param				pageSize			query		int			false "Page size"
//  @Success			200	{object}	responses.CorrectResponse{data=[]models.User,pagination=responses.Pagination}	"If the response is successful you will receive the requested user inside the data key"
//  @Failure			400	{object}	responses.FailureResponse	"If no token is provided the API will answer with a 400 code"
//  @Failure			403	{object}	responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed. the user has not verified their email yet or the user is not an admin"
//  @Failure			404	{object}	responses.FailureResponse	"If the if the user could not be found it will return a 404 code"
//  @Router 		/v1/user/ [get]
func GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	if user.Role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.NOT_AN_ADMIN,
			"message": "You do not have permission to access this resource",
		})
	}

	users, err := models.GetAllUsers(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_GET_USERS,
			"message": "Users could not be retrieved",
		})
	}

	page, pageSize := utils.Pagination(ctx)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.USERS_FOUND,
		"message": "Users found",
		"pagination": fiber.Map{
			"page":     page,
			"maxPages": math.Ceil(float64(models.CountUsers()) / float64(pageSize)),
			"pageSize": pageSize,
		},
		"data": users,
	})
}
