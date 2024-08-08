package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	"gorm.io/gorm"
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
	Type  string `gorm:"not null" json:"type" bson:"type"`
}

type Combat struct {
	Damage Damage `gorm:"type:jsonb;not null" json:"damage" bson:"damage"`
	Ac     uint8  `gorm:"not null" json:"ac" bson:"ac"`
}

type Item struct {
	gorm.Model

	ID          string            `gorm:"primary_key;not null" json:"id" bson:"id"`
	Name        pgtype.JSONBCodec `gorm:"type:jsonb;not null" json:"name" bson:"name"`
	Description pgtype.JSONBCodec `gorm:"type:jsonb;not null" json:"description" bson:"description"`
	Source      string            `gorm:"not null" json:"source" bson:"source"`
	Page        *uint16           `json:"page" bson:"page"`
	Tags        pq.StringArray    `gorm:"type:text;not null" json:"tags" bson:"tags"`
	Rarity      uint8             `gorm:"not null" json:"rarity" bson:"rarity"`
	Weight      float32           `gorm:"not null" json:"weight" bson:"weight"`
	Atunement   bool              `gorm:"not null" json:"atunement" bson:"atunement"`
	Cost        Cost              `gorm:"type:jsonb;not null" json:"cost" bson:"cost"`
	Contains    pq.StringArray    `gorm:"type:text;not null" json:"contains" bson:"contains"`
	Combat      Combat            `gorm:"type:jsonb;not null" json:"combat" bson:"combat"`
}

func (i *Item) Validate() error {
	return nil
}
