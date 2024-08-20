package event

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Event struct {
	ID         string     `json:"-"`
	App        string     `json:"app" validate:"required"`
	Type       string     `json:"type" validate:"required"`
	Time       time.Time  `json:"time" validate:"required"`
	Meta       Meta       `json:"meta" validate:"required"`
	Wallet     string     `json:"wallet" validate:"required"`
	Attributes Attributes `json:"attributes" validate:"required"`
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
