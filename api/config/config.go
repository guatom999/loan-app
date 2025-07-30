package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port string
	}

	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
)

func GetConfig() *Config {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		App: App{
			Port: os.Getenv("APP_PORT"),
		},
		Db: Db{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}
}
