package utils

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ExtractToken(ctx *fiber.Ctx) (string, error) {
	authorization := ctx.GetReqHeaders()["Authorization"]
	if len(authorization) < 1 {
		return "", errors.New("no authorization header provided");
	}

	authorizationContent := strings.Split(authorization[0], " ")
	if len(authorizationContent) != 2 {
		return "", errors.New("token not provided or malformed")
	}
	return authorizationContent[1], nil
}
