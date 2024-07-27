package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

func NewEnvConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
