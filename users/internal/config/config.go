package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server
	DB
}

type Server struct {
	Port     string
	Host     string
	LogLevel string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func New() *Config {
	if os.Getenv("LOAD_DOT_ENV") != "false" {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	conf := &Config{
		Server: Server{
			LogLevel: os.Getenv("SERVER_LOG_LEVEL"),
			Port:     os.Getenv("SERVER_PORT"),
			Host:     os.Getenv("SERVER_HOST"),
		},
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}

	slog.Debug(fmt.Sprintf("config: %+v", conf))
	return conf
}
