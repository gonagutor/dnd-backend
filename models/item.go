package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TranslatableField map[string]string

type Cost struct {
	Copper   uint8 `json:"copper" bson:"copper" default:"0"`
	Silver   uint8 `json:"silver" bson:"silver" default:"0"`
	Gold     uint8 `json:"gold" bson:"gold" default:"0"`
	Electrum uint8 `json:"electrum" bson:"electrum" default:"0"`
	Platinum uint8 `json:"platinum" bson:"platinum" default:"0"`
}

type Damage struct {
	Count uint8  `json:"count" bson:"count" validate:"required"`
	Faces uint8  `json:"faces" bson:"faces" validate:"required"`
	Range uint8  `json:"range" bson:"range" default:"0"`
	Type  string `json:"type" bson:"type" validate:"required"`
}

type Combat struct {
	Damage Damage `gorm:"type:jsonb" json:"damage" bson:"damage" validate:"optional,dive"`
	Ac     uint8  `json:"ac" bson:"ac" validate:"optional" default:"0"`
}

type Item struct {
	gorm.Model

	ID          string            `gorm:"primary_key" json:"id" bson:"id" validate:"required,unique"`
	Name        pgtype.JSONBCodec `gorm:"type:jsonb" json:"name" bson:"name" validate:"required,dive"`
	Description pgtype.JSONBCodec `gorm:"type:jsonb" json:"description" bson:"description" validate:"required,dive"`
	Source      string            `json:"source" bson:"source" validate:"required" default:"homebrew"`
	Page        *uint16           `json:"page" bson:"page" validate:"omitempty,min=1"`
	Tags        pq.StringArray    `gorm:"type:text" json:"tags" bson:"tags" validate:"required"`
	Rarity      uint8             `json:"rarity" bson:"rarity" validate:"required" default:"0"`
	Weight      float32           `json:"weight" bson:"weight" validate:"required" default:"0.0"`
	Atunement   bool              `json:"atunement" bson:"atunement" validate:"required" default:"false"`
	Cost        Cost              `gorm:"type:jsonb" json:"cost" bson:"cost" validate:"required,dive"`
	Contains    pq.StringArray    `gorm:"type:text" json:"contains" bson:"contains" validate:"optional,dive"`
	Combat      Combat            `gorm:"type:jsonb" json:"combat" bson:"combat" validate:"optional,dive"`
}

func (i *Item) Validate() error {
	return nil
}
