package database

import (
	"os"

	"github.com/bozkurtemre/backend-assesstment/src/frontend/internal/wallet"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := os.Getenv("DATA_CONNSTR")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&wallet.Wallet{}, wallet.Balance{})
	if err != nil {
		panic("failed to migrate database")
	}
}
