package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(ctx *fiber.Ctx) *gorm.DB {
	page, pageSize := Pagination(ctx)

	offset := (page - 1) * pageSize
	return PGConnection.Offset(offset).Limit(pageSize)
}
