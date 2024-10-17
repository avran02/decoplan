package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ExternalServices
	Server Server
}

type ExternalServices struct {
	AuthServiceUrl string
}

type Server struct {
	LogLevel string
	Port     string
	Host     string
}

func New() *Config {
	if os.Getenv("LOAD_DOT_ENV") != "false" {
		slog.Info("Loading .env file")
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	config := &Config{
		ExternalServices: ExternalServices{
			AuthServiceUrl: os.Getenv("AUTH_SERVER_URL"),
		},
		Server: Server{
			LogLevel: os.Getenv("SERVER_LOG_LEVEL"),
			Port:     os.Getenv("SERVER_PORT"),
			Host:     os.Getenv("SERVER_HOST"),
		},
	}
	slog.Debug(fmt.Sprintf("config: %+v", config))

	return config
}
