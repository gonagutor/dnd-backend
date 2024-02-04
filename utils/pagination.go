package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Pagination(ctx *fiber.Ctx) (page int, pageSize int) {
	pageQuery := ctx.Query("page")
	page, pageError := strconv.Atoi(pageQuery)
	if pageError != nil || page <= 0 {
		page = 1
	}

	pageSizeQuery := ctx.Query("page_size")
	pageSize, pageSizeError := strconv.Atoi(pageSizeQuery)
	if pageSizeError != nil || pageSize <= 0 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	return page, pageSize
}
