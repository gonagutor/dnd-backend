package models

import (
	"cardando/backend/utils"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uuid.UUID      `gorm:"type:uuid;primary_key;"`
	Name      string         `gorm:"not null"`
	Surname   string         `gorm:"not null"`
	Username  string         `gorm:"not null"`
	Email     string         `gorm:"not null"`
	Password  string         `gorm:"not null"`
	IsActive  bool           `gorm:"not null;default:false"`
	Roles     pq.StringArray `gorm:"type:varchar(64)[];not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)

	if user.Roles == nil || len(user.Roles) < 1 {
		user.Roles = pq.StringArray{"user"}
	}
	return nil
}

func FindUserByUsername(username string) (*User, error) {
	ret := &User{Username: username}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
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
