package models

import (
	"revosearch/backend/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email          string    `gorm:"type:varchar(256);not null"`
	Name           string    `gorm:"type:varchar(32);not null"`
	Surname        string    `gorm:"type:varchar(64);not null"`
	Password       string    `gorm:"type:varchar(128);not null"`
	Role           string    `gorm:"type:varchar(16);not null;default:user"`
	ProfilePicture string    `gorm:"default:null"`
	IsActive       bool      `gorm:"not null;default:false"`
	DeletedAt      *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)
	return nil
}

func FindUserByEmail(email string) (*User, error) {
	ret := &User{Email: email}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func FindUserByID(id uuid.UUID) (*User, error) {
	ret := &User{ID: id}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
