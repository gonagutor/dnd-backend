package models

import (
	utils_constants "revosearch/backend/constants/utils"
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
	RefreshKey     string    `gorm:"type:varchar(16);not null"`
	IsActive       bool      `gorm:"not null;default:false"`
	DeletedAt      *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), utils_constants.PASSWORD_COST)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)
	user.RefreshKey = utils.GenerateRandomCode(15)
	return nil
}

func FindUserByEmail(email string) (*User, error) {
	ret := &User{Email: email}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func FindUserByID(id string) (*User, error) {
	idParsed, parseError := uuid.Parse(id)
	if parseError != nil {
		return nil, parseError
	}

	ret := &User{ID: idParsed}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) CheckKey(key string) error {
	return bcrypt.CompareHashAndPassword([]byte(key), []byte(user.ID.String()+user.RefreshKey))
}
