package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
	ServerAddress string
}

func LoadConfig() (Config, error) {
	var cfg Config

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Load configuration values from environment variables
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")

	// Validate required configuration values
	if cfg.ServerAddress == "" {
		return cfg, fmt.Errorf("SERVER_ADDRESS is not set")
	}
	if cfg.DBHost == "" {
		return cfg, fmt.Errorf("DB_HOST is not set")
	}
	if cfg.DBPort == "" {
		return cfg, fmt.Errorf("DB_PORT is not set")
	}
	if cfg.DBUser == "" {
		return cfg, fmt.Errorf("DB_USER is not set")
	}
	if cfg.DBPassword == "" {
		return cfg, fmt.Errorf("DB_PASSWORD is not set")
	}
	if cfg.DBName == "" {
		return cfg, fmt.Errorf("DB_NAME is not set")
	}
	if cfg.JWTSecret == "" {
		return cfg, fmt.Errorf("JWT_SECRET is not set")
	}

	return cfg, nil
}
