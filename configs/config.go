package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabasePassword string
	DatabaseUsername string

	Port string
}

var (
	Env = GetConfig()
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Application is running without environment variables")
	}
}

func GetConfig() Config {
	var config Config

	LoadEnv()
	config.DatabaseHost = os.Getenv("DB_HOST")
	config.DatabasePort = os.Getenv("DB_PORT")
	config.DatabaseName = os.Getenv("DB_NAME")
	config.DatabasePassword = os.Getenv("DB_PASSWORD")
	config.DatabaseUsername = os.Getenv("DB_USERNAME")

	config.Port = os.Getenv("APP_PORT")

	return config
}

func GetIntFromEnv(key string) int {
	env := os.Getenv(key)
	intEnv, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}
	return intEnv
}
