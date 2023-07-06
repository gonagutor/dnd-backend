package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlatSearch struct {
	gorm.Model

	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null"`
	Owner       User      `gorm:"foreignKey:OwnerID"`
	Name        string    `gorm:"default:null"`
	Description string    `gorm:"default:null"`
	Visiblity   string    `gorm:"default:private;not null"`
	Users       []User

	DeletedAt *time.Time `gorm:"default:null"`
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"not null;default:current_timestamp"`
}
