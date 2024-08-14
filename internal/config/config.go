package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DOToken string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DOToken: os.Getenv("DO_TOKEN"),
	}, nil
}