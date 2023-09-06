package models

import (
	"dnd/backend/utils"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model

	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	ClassID      uuid.UUID `gorm:"type:uuid;not null"`
	RaceID       uuid.UUID `gorm:"type:uuid;not null"`
	BackgroundID uuid.UUID `gorm:"type:uuid;not null"`
	Name         string    `gorm:"not null"`

	XP                 int `gorm:"not null"`
	AC                 int `gorm:"not null"`
	HP                 int `gorm:"not null"`
	TemporaryHP        int `gorm:"not null"`
	DeathSaveSuccesses int `gorm:"not null"`
	DeathSaveFailures  int `gorm:"not null"`

	Inspiration      int `gorm:"not null;default:0"`
	ProficiencyBonus int `gorm:"not null;default:2"`

	Strength     int `gorm:"not null;default:0"`
	Dexterity    int `gorm:"not null;default:0"`
	Constitution int `gorm:"not null;default:0"`
	Intelligence int `gorm:"not null;default:0"`
	Wisdom       int `gorm:"not null;default:0"`
	Charisma     int `gorm:"not null;default:0"`

	Alignment         string
	PersonalityTraits string
	Ideals            string
	Bonds             string
	Flaws             string

	CopperCoins   int `gorm:"not null"`
	SilverCoins   int `gorm:"not null"`
	ElectrumCoins int `gorm:"not null"`
	GoldCoins     int `gorm:"not null"`
	PlatinumCoins int `gorm:"not null"`

	Equipment    pq.StringArray `gorm:"type:text[];not null"` // These are UUID arrays
	Feats        pq.StringArray `gorm:"type:text[];not null"` // These are UUID arrays
	ActiveSpells pq.StringArray `gorm:"type:text[];not null"` // These are UUID arrays

	Languages          pq.StringArray `gorm:"type:text[];not null"`
	OtherProficiencies pq.StringArray `gorm:"type:text[];not null"`

	Acrobatics     bool `gorm:"not null;default:false"`
	AnimalHandling bool `gorm:"not null;default:false"`
	Arcana         bool `gorm:"not null;default:false"`
	Athletics      bool `gorm:"not null;default:false"`
	Deception      bool `gorm:"not null;default:false"`
	History        bool `gorm:"not null;default:false"`
	Insight        bool `gorm:"not null;default:false"`
	Intimidation   bool `gorm:"not null;default:false"`
	Investigation  bool `gorm:"not null;default:false"`
	Medicine       bool `gorm:"not null;default:false"`
	Nature         bool `gorm:"not null;default:false"`
	Perception     bool `gorm:"not null;default:false"`
	Performance    bool `gorm:"not null;default:false"`
	Persuasion     bool `gorm:"not null;default:false"`
	Religion       bool `gorm:"not null;default:false"`
	SleightOfHand  bool `gorm:"not null;default:false"`
	Steath         bool `gorm:"not null;default:false"`
	Survival       bool `gorm:"not null;default:false"`

	DeletedAt *time.Time `gorm:"default:null"`
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"not null;default:current_timestamp"`
}

func RestoreDeletedCharacter(id string) error {
	idParsed, parseError := uuid.Parse(id)
	if parseError != nil {
		return parseError
	}

	return utils.PGConnection.Model(&Character{}).Where("id", idParsed).Update("deleted_at", nil).Error
}

func FindCharacterByID(id string) (*Character, error) {
	idParsed, parseError := uuid.Parse(id)
	if parseError != nil {
		return nil, parseError
	}

	ret := &Character{ID: idParsed}
	if err := utils.PGConnection.First(ret, ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}
