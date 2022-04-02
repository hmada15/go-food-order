package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Env   string
	DEBUG string
	MySQL mySQLConfig
	Port  string
}

type mySQLConfig struct {
	Username     string
	Password     string
	DatabaseName string
	Host         string
	Port         string
}

func fallback(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func NewConfig() config {
	Config := config{}
	Config.Env = fallback("ENV", "false")
	Config.Port = fallback("PORT", "1323")
	Config.DEBUG = fallback("DEBUG", "false")
	Config.MySQL.DatabaseName = fallback("DB_DATABASE", "go-")
	Config.MySQL.Username = fallback("DB_USERNAME", "root")
	Config.MySQL.Password = fallback("DB_PASSWORD", "")
	Config.MySQL.Host = fallback("DB_HOST", "127.0.0.1")
	Config.MySQL.Port = fallback("DB_PORT", "3306")

	return Config
}
