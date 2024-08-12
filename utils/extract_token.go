package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ExtractToken(ctx *fiber.Ctx) (string, error) {
	authorization := ctx.GetReqHeaders()["Authorization"][0]
	authorizationContent := strings.Split(authorization, " ")
	if len(authorizationContent) != 2 {
		return "", errors.New("token not provided or malformed")
	}
	return authorizationContent[1], nil
}
