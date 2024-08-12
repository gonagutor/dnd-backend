package v1_item_handler

import (
	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
)

type ItemBody struct {
	ID          string            `json:"id"`
	Name        pgtype.JSONBCodec `json:"name"`
	Description pgtype.JSONBCodec `json:"descriptom"`

	Source string  `json:"source"`
	Page   *uint16 `json:"page"`

	Tags      pq.StringArray `json:"tags"`
	Rarity    uint8          `json:"rarity"`
	Weight    float32        `json:"weight"`
	Atunement bool           `json:"atunement"`

	Cost     models.Cost    `json:"cost"`
	Contains pq.StringArray `json:"contains"`
	Combat   models.Combat  `json:"combat"`
}

func Create(ctx *fiber.Ctx) error {
	item, err := parseAndValidate(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.INVALID_DATA,
			"message": err,
		})
	}

	creationError := utils.PGConnection.Create(item).Error
	if creationError != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_CREATE_ITEM,
			"message": "Could not create item",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.ITEM_CREATED,
		"message": "Item created successfully",
	})
}

func parseAndValidate(ctx *fiber.Ctx) (*models.Item, error) {
	itemBody := new(ItemBody)
	err := ctx.BodyParser(itemBody)
	if err != nil {
		return nil, errors.New("Error parsing the body")
	}

	user := protected.GetUserFromContext(ctx)

	item := &models.Item{
		ID:          itemBody.ID,
		Name:        itemBody.Name,
		Description: itemBody.Description,
		Source:      itemBody.Source,
		Page:        itemBody.Page,
		Tags:        itemBody.Tags,
		Rarity:      itemBody.Rarity,
		Weight:      itemBody.Weight,
		Atunement:   itemBody.Atunement,
		Cost:        itemBody.Cost,
		Contains:    itemBody.Contains,
		Combat:      itemBody.Combat,
		User:        &user.ID,
	}
	err = item.Validate()
	if err != nil {
		return nil, err
	}

	return item, nil
}
