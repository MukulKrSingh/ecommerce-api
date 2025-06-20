package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()
	err = godotenv.Load(filepath.Join("/Users/mukulkumarsingh/workspace/go_space/ecommerce-api/", ".env"))
	if err != nil {
		return AppConfig{}, fmt.Errorf("error loading .env file: %w", err)
	}

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("SERVER_PORT")
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	AppSecret := os.Getenv("APP_SECRET")
	if AppSecret == "" {
		return AppConfig{}, fmt.Errorf("APP_SECRET is not set in the environment variables")
	} else {
		log.Println("APP_SECRET is set in the environment variables")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: AppSecret}, nil
}
