package wallet

import "time"

type Wallet struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserId    string    `gorm:"not null" json:"-"`
	Balances  []Balance `gorm:"foreignKey:WalletID" json:"balances"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Balance struct {
	ID        string    `gorm:"primaryKey" json:"-"`
	WalletID  string    `gorm:"not null" json:"-"`
	Currency  string    `gorm:"not null" json:"currency"`
	Amount    string    `gorm:"type:decimal(10,2); not null" json:"amount"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
