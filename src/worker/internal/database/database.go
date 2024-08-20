package database

import (
	"github.com/bozkurtemre/backend-assesstment/src/worker/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := config.Config("DATA_CONNSTR")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
