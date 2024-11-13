package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	BaseURL    string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	JwtSecret  string
	Mode       string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		envConfig = &Config{
			Port:       os.Getenv("APP_PORT"),
			BaseURL:    os.Getenv("BASE_URL"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBHost:     os.Getenv("DB_HOST"),
			JwtSecret:  os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			Mode:       os.Getenv("MODE"),
		}
	})

	return envConfig
}
