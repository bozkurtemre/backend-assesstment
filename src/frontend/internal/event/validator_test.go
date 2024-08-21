package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEvent(t *testing.T) {
	tests := []struct {
		name    string
		event   Event
		wantErr error
	}{
		{
			name: "valid balance increase",
			event: Event{
				Type: EventTypeBalanceIncrease,
				Attributes: Attributes{
					Amount:   "33.20",
					Currency: "TRY",
				},
			},
			wantErr: nil,
		},
		{
			name: "valid balance decrease",
			event: Event{
				Type: EventTypeBalanceDecrease,
				Attributes: Attributes{
					Amount:   "3.10",
					Currency: "TRY",
				},
			},
			wantErr: nil,
		},
		{
			name: "invalid event type",
			event: Event{
				Type: "INVALID_TYPE",
				Attributes: Attributes{
					Amount:   "33.20",
					Currency: "USD",
				},
			},
			wantErr: ErrInvalidEventType,
		},
		{
			name: "invalid amount",
			event: Event{
				Type: EventTypeBalanceIncrease,
				Attributes: Attributes{
					Amount:   "invalid_amount",
					Currency: "TRY",
				},
			},
			wantErr: ErrInvalidAmount,
		},
		{
			name: "invalid currency",
			event: Event{
				Type: EventTypeBalanceIncrease,
				Attributes: Attributes{
					Amount:   "33.20",
					Currency: "TR",
				},
			},
			wantErr: ErrInvalidCurrency,
		},
		{
			name: "invalid currency with non-letter characters",
			event: Event{
				Type: EventTypeBalanceIncrease,
				Attributes: Attributes{
					Amount:   "3.10",
					Currency: "US1",
				},
			},
			wantErr: ErrInvalidCurrency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEvent(tt.event)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
