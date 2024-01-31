package models

type TranslatableField map[string]string

type HitDice struct {
	Count uint8  `json:"count" bson:"count" validate:"min=1,max=255,required"`
	Faces uint16 `json:"faces" bson:"faces" validate:"min=1,max=65535,required"`
}

type SkillChoice struct {
	Count uint8    `json:"count" bson:"count" validate:"required"`
	From  []string `json:"from" bson:"from" validate:"required"`
}

type SkillsProficiencies struct {
	Choose SkillChoice `json:"choose" bson:"choose" validate:"required,dive"`
	Forced []string    `json:"forced" bson:"forced" validate:"required"`
}

type Proficiencies struct {
	SavingThrows []string            `json:"savingThrows" bson:"savingThrows" validate:"required"`
	Skills       SkillsProficiencies `json:"skills" bson:"skills" validate:"required,dive"`
	Weapons      []string            `json:"weapons" bson:"weapons" validate:"required"`
	Armor        []string            `json:"armor" bson:"armor" validate:"required"`
}

type Equipment struct {
	AdditionalFromBackground bool `json:"additionalFromBackground" bson:"additionalFromBackground" validate:"required"`
}

type ClassData struct {
	Name   TranslatableField `json:"name" bson:"name" validate:"required,dive"`
	Source *string           `json:"source" bson:"source"`
	Page   *uint16           `json:"page" bson:"page" validate:"omitempty,min=1"`
	Srd    *bool             `json:"srd" bson:"srd"`

	HitDice       HitDice       `json:"hitDice" bson:"hitDice" validate:"required,dive"`
	Proficiencies Proficiencies `json:"proficiencies" bson:"proficiencies" validate:"required,dive"`
	Equipment     Equipment     `json:"equipment" bson:"equipment" validate:"required,dive"`
}

type Class struct {
	Class []ClassData `json:"class" bson:"class" validate:"required,dive,max=32,min=1"`
}