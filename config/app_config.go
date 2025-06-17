package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load(".env")

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load("../.env")
	}

	httpPort := os.Getenv("SERVER_PORT")
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	return AppConfig{ServerPort: httpPort, Dsn: Dsn}, nil
}
