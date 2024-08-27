package models

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"dnd/backend/utils"
)

type TranslatableField map[string]string

type Cost struct {
	Copper   uint32 `gorm:"not null" json:"copper" bson:"copper"`
	Electrum uint32 `gorm:"not null" json:"electrum" bson:"electrum"`
	Gold     uint32 `gorm:"not null" json:"gold" bson:"gold"`
	Silver   uint32 `gorm:"not null" json:"silver" bson:"silver"`
	Platinum uint32 `gorm:"not null" json:"platinum" bson:"platinum"`
}

type Damage struct {
	Count uint8  `gorm:"not null" json:"count" bson:"count"`
	Faces uint8  `gorm:"not null" json:"faces" bson:"faces"`
	Range uint16 `gorm:"not null" json:"range" bson:"range"`
	Type  uint8  `gorm:"not null" json:"type" bson:"type"`
}

type Combat struct {
	Damage Damage `gorm:"type:jsonb;not null" json:"damage" bson:"damage"`
	Ac     uint8  `gorm:"not null" json:"ac" bson:"ac"`
}

type Item struct {
	gorm.Model `swaggerignore:"true"`

	ID          string            `gorm:"primary_key;not null" json:"id" bson:"id"`
	Name        pgtype.JSONBCodec `gorm:"type:jsonb;not null" json:"name" bson:"name" swaggertype:"object"`
	Description pgtype.JSONBCodec `gorm:"type:jsonb;not null" json:"description" bson:"description" swaggertype:"object"`

	Source string  `gorm:"not null" json:"source" bson:"source"`
	Page   *uint16 `json:"page" bson:"page"`

	Tags      pq.StringArray `gorm:"type:text;not null" json:"tags" bson:"tags" swaggertype:"array,string"`
	Rarity    uint8          `gorm:"not null" json:"rarity" bson:"rarity"`
	Weight    float32        `gorm:"not null" json:"weight" bson:"weight"`
	Atunement bool           `gorm:"not null" json:"atunement" bson:"atunement"`

	Cost     Cost           `gorm:"type:jsonb;not null" json:"cost" bson:"cost"`
	Contains pq.StringArray `gorm:"type:text;not null" json:"contains" bson:"contains" swaggertype:"array,string"`
	Combat   Combat         `gorm:"type:jsonb;not null" json:"combat" bson:"combat"`

	User *uuid.UUID `gorm:"type:text" json:"user" bson:"user"`
}

func FindItemByID(id string) (*Item, error) {
	ret := &Item{ID: id}
	err := utils.PGConnection.First(ret, ret).Error
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func GetAllItems(ctx *fiber.Ctx) ([]Item, error) {
	key := ctx.Query("key")
	if key == "" {
		key = "id"
	}

	sortOrder := ctx.Query("sortOrder")
	if sortOrder == "" {
		sortOrder = "DESC"
	}

	order := key + " " + sortOrder
	var items []Item

	err := utils.Paginate(ctx).Order(order).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemsFromUser(user uuid.UUID) ([]Item, error) {
	var items []Item
	err := utils.PGConnection.Where("user = ? OR user IS NULL", user).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func CountItems() int64 {
	var count int64
	utils.PGConnection.Model(&Item{}).Count(&count)
	return count
}

func (i *Item) Validate() error {
	item, _ := FindItemByID(i.ID)
	if item != nil {
		return errors.New("item ID already exists")
	}

	if *i.Page == 0 {
		i.Page = nil
	}

	for index, tag := range i.Tags {
		if strings.Contains(tag, "#"+i.ID) {
			return errors.New("item can not refer to itself")
		}

		if containsSpecialCharOrNumber(tag) {
			return errors.New("special characters or numbers not allowed in tag")
		}

		i.Tags[index] = strings.ToLower(tag)
		if !isKebabCase(strings.ToLower(tag)) {
			return errors.New("tag must be in kebab-case/dash-case")
		}
	}

	if i.Combat.Damage.Count == 0 {
		return errors.New("'count' stat in Damage is required")
	}

	if i.Combat.Damage.Faces == 0 {
		return errors.New("'Faces' stat in Damage is required")
	}

	if i.Combat.Damage.Type == 0 {
		return errors.New("'Type' stat in Damage is required")
	}

	return nil
}

func containsSpecialCharOrNumber(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && char != '#' {
			return true
		}
	}
	return false
}

func isKebabCase(s string) bool {
	kebabCasePattern := `^#?[a-z]+(?:-[a-z]+)*$`

	re := regexp.MustCompile(kebabCasePattern)
	return re.MatchString(s)
}
