package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbURL	 string
	DbDriver string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbURL := os.Getenv("DB_URL")

	return &Config{
		DbURL:    dbURL,
		DbDriver: dbDriver,
	}, nil
}
