package models

import (
	utils_constants "dnd/backend/constants/utils"
	"dnd/backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email          string    `gorm:"type:varchar(256);not null" json:"email"`
	Name           string    `gorm:"type:varchar(32);not null" json:"name"`
	Surname        string    `gorm:"type:varchar(64);not null" json:"surname"`
	Password       string    `gorm:"type:varchar(128);not null"`
	Role           string    `gorm:"type:varchar(16);not null;default:user" json:"role"`
	ProfilePicture string    `gorm:"default:null" json:"profilePicture"`
	RefreshKey     string    `gorm:"type:varchar(16);not null" json:"refreshKey"`
	IsActive       bool      `gorm:"not null;default:false" json:"isActive"`

	DeletedAt *time.Time `gorm:"default:null" json:"deletedAt"`
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"not null;default:current_timestamp" json:"updatedAt"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), utils_constants.PASSWORD_COST)
	if err != nil {
		return err
	}
	user.Password = string(passwordHashed)
	user.RefreshKey = utils.GenerateRandomCode(utils_constants.REFRESH_KEY_LENGTH)
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

func GetAllUsers(ctx *fiber.Ctx) ([]User, error) {
	key := ctx.Query("key")
	if key == "" {
		key = "created_at"
	}

	sortOrder := ctx.Query("sortOrder")
	if sortOrder == "" {
		sortOrder = "DESC"
	}

	order := key + " " + sortOrder
	var users []User

	err := utils.Paginate(ctx).Omit("password", "refresh_key").Order(order).Find(&users)
	if err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

func CountUsers() int64 {
	var count int64
	utils.PGConnection.Model(&User{}).Count(&count)
	return count
}

func UserExistsByID(id string) bool {
	idParsed, parseError := uuid.Parse(id)
	if parseError != nil {
		return false
	}

	var count int64 = 0
	utils.PGConnection.First(&User{ID: idParsed}).Count(&count)
	if count == 0 {
		return false
	}

	return true
}
