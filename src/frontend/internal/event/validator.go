package event

import (
	"errors"
	"strings"
	"unicode"
)

const (
	EventTypeBalanceIncrease = "BALANCE_INCREASE"
	EventTypeBalanceDecrease = "BALANCE_DECREASE"
)

var (
	ErrInvalidEventType = errors.New("invalid event type")
	ErrInvalidAmount    = errors.New("invalid amount")
	ErrInvalidCurrency  = errors.New("invalid currency")
)

func ValidateEvent(event Event) error {
	if event.Type != EventTypeBalanceIncrease && event.Type != EventTypeBalanceDecrease {
		return ErrInvalidEventType
	}

	if !isValidAmount(event.Attributes.Amount) {
		return ErrInvalidAmount
	}

	if !isValidCurrency(event.Attributes.Currency) {
		return ErrInvalidCurrency
	}

	return nil
}

func isValidAmount(amount string) bool {
	return strings.Contains(amount, ".")
}

func isValidCurrency(currency string) bool {
	if len(currency) != 3 {
		return false
	}

	for _, char := range currency {
		if !unicode.IsLetter(char) {
			return false
		}
	}

	return true
}
