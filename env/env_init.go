package env

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	// DB_URL      string
	DB_USER     string
	DB_PASSWORD string
	DB_ADDR     string
	DB_NAME     string
	PORT        string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB_URL = getEnvWithDefault("DB_URL", "postgres://ductt:123@localhost:5432/meetmeup?sslmode=disable")

	DB_USER = getEnvWithDefault("DB_USER", "ductt")
	DB_PASSWORD = getEnvWithDefault("DB_PASSWORD", "123")
	DB_ADDR = getEnvWithDefault("DB_ADDR", "localhost:5432")
	DB_NAME = getEnvWithDefault("DB_NAME", "meetmeup")

	PORT = getEnvWithDefault("PORT", "8080")

}
