package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server           Server
	ExternalServices ExternalServices
}

type Server struct {
	LogLevel string
	Port     string
	Host     string
}

type ExternalServices struct {
	AuthURL    string
	StorageURL string
	UsersURL   string
}

func New() *Config {
	if os.Getenv("LOAD_DOT_ENV") != "false" {
		slog.Info("Loading .env file")
		if err := godotenv.Load(); err != nil {
			slog.Error(err.Error())
		}
	}

	config := &Config{
		Server: Server{
			LogLevel: os.Getenv("SERVER_LOG_LEVEL"),
			Port:     os.Getenv("SERVER_PORT"),
			Host:     os.Getenv("SERVER_HOST"),
		},
		ExternalServices: ExternalServices{
			AuthURL:    os.Getenv("AUTH_SERVER_URL"),
			StorageURL: os.Getenv("STORAGE_SERVER_URL"),
			UsersURL:   os.Getenv("USERS_SERVER_URL"),
		},
	}

	return config
}
