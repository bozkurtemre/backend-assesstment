package event

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Event struct {
	ID         string     `gorm:"primaryKey" json:"-"`
	App        string     `gorm:"not null" json:"app" validate:"required"`
	Type       string     `gorm:"not null" json:"type" validate:"required"`
	Time       time.Time  `gorm:"not null" json:"time" validate:"required"`
	Meta       Meta       `gorm:"type:json;not null" json:"meta" validate:"required"`
	Wallet     string     `gorm:"not null" json:"wallet" validate:"required"`
	Attributes Attributes `gorm:"type:json;not null" json:"attributes" validate:"required"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
}

type Meta struct {
	User string `json:"user" validate:"required"`
}

type Attributes struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type Events struct {
	Events []Event `json:"events" validate:"required,dive,required"`
}

var Validator = validator.New()
