package wallet

import (
	"errors"
	"log"
)

type Service struct {
	walletRepository WalletRepository
}

type WalletService interface {
	UpdateUserWalletBalance(walletID string, currency string, amount float64) error
}

func NewWalletService(walletRepository WalletRepository) WalletService {
	return &Service{walletRepository: walletRepository}
}

func (s *Service) UpdateUserWalletBalance(walletID string, currency string, amount float64) error {
	wallet, err := s.walletRepository.FindByID(walletID)
	if err != nil {
		return err
	}

	if wallet == nil {
		log.Fatalf("wallet not found: %v", walletID)
	}

	balance, err := s.walletRepository.FindBalanceByWalletIDAndCurrency(walletID, currency)
	if err != nil {
		return err
	}

	newBalance := balance.Amount + amount
	if newBalance < 0 {
		return errors.New("insufficient balance")
	}
	balance.Amount = newBalance

	return s.walletRepository.UpdateBalance(balance)
}
