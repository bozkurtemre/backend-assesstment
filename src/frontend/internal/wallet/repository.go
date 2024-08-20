package wallet

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type WalletRepository interface {
	FindAll() (*[]Wallet, error)
}

func NewWalletRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindAll() (*[]Wallet, error) {
	var wallets []Wallet
	r.db.Preload("Balances").Find(&wallets)

	return &wallets, nil
}
