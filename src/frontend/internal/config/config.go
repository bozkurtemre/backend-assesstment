package config

import (
	"os"

	_ "github.com/joho/godotenv"
)

func Config(key string) string {
	return os.Getenv(key)
}
