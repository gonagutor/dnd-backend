package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flat struct {
	gorm.Model

	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FlatSearchID uuid.UUID  `gorm:"type:uuid;not null"`
	FlatSearch   FlatSearch `gorm:"foreignKey:FlatSearchID"`
	Name         string     `gorm:"default:null"`
	Description  string     `gorm:"default:null"`

	DeletedAt *time.Time `gorm:"default:null"`
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"not null;default:current_timestamp"`
}
