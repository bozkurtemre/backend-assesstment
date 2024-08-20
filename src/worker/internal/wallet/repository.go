package wallet

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

type WalletRepository interface {
	FindByID(id string) (*Wallet, error)
	FindBalanceByWalletIDAndCurrency(walletID string, currency string) (*Balance, error)
	UpdateBalance(balance *Balance) error
}

func NewWalletRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindByID(id string) (*Wallet, error) {
	var wallet Wallet
	if err := r.DB.Where("id = ?", id).Preload("Balances").First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *Repository) FindBalanceByWalletIDAndCurrency(walletID string, currency string) (*Balance, error) {
	var balance Balance
	if err := r.DB.Where("wallet_id = ? AND currency = ?", walletID, currency).First(&balance).Error; err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *Repository) UpdateBalance(balance *Balance) error {
	return r.DB.Save(balance).Error
}
