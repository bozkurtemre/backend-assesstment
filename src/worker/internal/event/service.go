package event

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/logger"
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/wallet"
)

type Service struct {
	walletService  wallet.WalletService
	loggerInstance *logger.Logger
}

type EventService interface {
	ProcessEvent(eventData []byte) error
}

func NewEventService(walletService wallet.WalletService, loggerInstance *logger.Logger) EventService {
	return &Service{walletService: walletService, loggerInstance: loggerInstance}
}

func (s *Service) ProcessEvent(eventData []byte) error {
	var event Event

	if err := json.Unmarshal(eventData, &event); err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(event.Attributes.Amount, 64)
	if err != nil {
		return err
	}

	err = s.walletService.UpdateUserWalletBalance(event.Wallet, event.Attributes.Currency, s.detectAmount(amount, event.Type))
	if err != nil {
		return err
	}

	err = s.loggerInstance.Log(eventData)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Service) detectAmount(amount float64, event string) float64 {
	switch event {
	case BalanceIncrease:
		return amount
	case BalanceDecrease:
		return -amount
	default:
		return 0
	}
}
