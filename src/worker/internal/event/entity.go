package event

import (
	"time"
)

const (
	BalanceIncrease = "BALANCE_INCREASE"
	BalanceDecrease = "BALANCE_DECREASE"
)

type Event struct {
	ID         string     `json:"-"`
	App        string     `json:"app"`
	Type       string     `json:"type"`
	Time       time.Time  `json:"time"`
	Meta       Meta       `json:"meta"`
	Wallet     string     `json:"wallet"`
	Attributes Attributes `json:"attributes"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
}

type Meta struct {
	User string `json:"user"`
}

type Attributes struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
