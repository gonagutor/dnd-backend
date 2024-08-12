package v1_item_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"

	"dnd/backend/constants/http_codes"
	"dnd/backend/errors/http_errors"
	"dnd/backend/middleware/protected"
	"dnd/backend/models"
	"dnd/backend/utils"

)

type EditCost struct {
	Copper   *uint32 `json:"copper"`
	Electrum *uint32 `json:"electrum"`
	Gold     *uint32 `json:"gold"`
	Silver   *uint32 `json:"silver"`
	Platinum *uint32 `json:"platinum"`
}

type EditDamage struct {
	Count *uint8  `json:"count"`
	Faces *uint8  `json:"faces"`
	Range *uint16 `json:"range"`
	Type  *uint8  `json:"type"`
}

type EditCombat struct {
	Damage *EditDamage `json:"damage"`
	Ac     *uint8      `json:"ac"`
}

type EditItemBody struct {
	ID          *string            `json:"id"`
	Name        *pgtype.JSONBCodec `json:"name"`
	Description *pgtype.JSONBCodec `json:"descriptom"`

	Source *string `json:"source"`
	Page   *uint16 `json:"page"`

	Tags      *pq.StringArray `json:"tags"`
	Rarity    *uint8          `json:"rarity"`
	Weight    *float32        `json:"weight"`
	Atunement *bool           `json:"atunement"`

	Cost     *EditCost       `json:"cost"`
	Contains *pq.StringArray `json:"contains"`
	Combat   *EditCombat     `json:"combat"`
}

// @Tags Item
// @Description Edit an item
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token with Bearer prefix"
// @Param item_id path string true "The id of the item you want to edit"
// @Param body body EditItemBody true "The body of the item you want to edit"
// @Success 200 {object} responses.CorrectResponse "If the response is successful you will receive a simple code and message indicating that the item has been edited"
// @Failure 400 {object} responses.FailureResponse "If the request is malformed or the data is invalid"
// @Failure 400 {object} responses.FailureResponse "If no token is provided the API will answer with a 400 code"
// @Failure 403 {object} responses.FailureResponse "The API can answer with a 403 if the token is invalid/malformed, the user has not verified their email yet or (if the user is trying to edit a public item) the user is not an admin"
// @Failure 404 {object} responses.FailureResponse "If the item is not found"
// @Failure 500 {object} responses.FailureResponse "If the server fails to edit the item it will answer with a 500 code. Please report this error if you encounter it in production"
// @Router /v1/item/edit/{item_id} [put]
func Edit(ctx *fiber.Ctx) error {
	editItem := new(EditItemBody)
	err := ctx.BodyParser(editItem)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.BAD_REQUEST,
			"message": "Body could not be parsed",
		})
	}

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
			"message": "You can only edit your own items",
		})
	}

	err = editItemAndValidate(itemUrl, editItem)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   http_errors.INVALID_DATA,
			"message": err,
		})
	}

	err = utils.PGConnection.Save(&itemUrl).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   http_errors.COULD_NOT_EDIT_USER,
			"message": "Item could not be edited",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    http_codes.ITEM_EDITED,
		"message": "Item edited",
	})
}

func editItemAndValidate(item *models.Item, editItem *EditItemBody) error {
	if editItem.Name != nil {
		item.Name = *editItem.Name
	}
	if editItem.Description != nil {
		item.Description = *editItem.Description
	}

	if editItem.Source != nil {
		item.Source = *editItem.Source
	}
	if editItem.Page != nil {
		item.Page = editItem.Page
	}

	if editItem.Tags != nil {
		item.Tags = *editItem.Tags
	}
	if editItem.Rarity != nil {
		item.Rarity = *editItem.Rarity
	}
	if editItem.Weight != nil {
		item.Weight = *editItem.Weight
	}
	if editItem.Atunement != nil {
		item.Atunement = *editItem.Atunement
	}

	if editItem.Cost != nil {
		if editItem.Cost.Copper != nil {
			item.Cost.Copper = *editItem.Cost.Copper
		}
		if editItem.Cost.Electrum != nil {
			item.Cost.Electrum = *editItem.Cost.Electrum
		}
		if editItem.Cost.Gold != nil {
			item.Cost.Gold = *editItem.Cost.Gold
		}
		if editItem.Cost.Silver != nil {
			item.Cost.Silver = *editItem.Cost.Silver
		}
		if editItem.Cost.Platinum != nil {
			item.Cost.Platinum = *editItem.Cost.Platinum
		}
	}

	if editItem.Contains != nil {
		item.Contains = *editItem.Contains
	}

	if editItem.Combat.Damage != nil {
		if editItem.Combat.Damage.Count != nil {
			item.Combat.Damage.Count = *editItem.Combat.Damage.Count
		}
		if editItem.Combat.Damage.Faces != nil {
			item.Combat.Damage.Faces = *editItem.Combat.Damage.Faces
		}
		if editItem.Combat.Damage.Range != nil {
			item.Combat.Damage.Range = *editItem.Combat.Damage.Range
		}
		if editItem.Combat.Damage.Type != nil {
			item.Combat.Damage.Type = *editItem.Combat.Damage.Type
		}
	}
	if editItem.Combat.Ac != nil {
		item.Combat.Ac = *editItem.Combat.Ac
	}

	return item.Validate()
}