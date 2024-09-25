package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Minio  Minio
	Server Server
}

type Minio struct {
	Endpoint  string
	AccessKey string
	SecretKey string
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
		Minio: Minio{
			Endpoint:  os.Getenv("MINIO_ENDPOINT"),
			AccessKey: os.Getenv("MINIO_ACCESS_KEY"),
			SecretKey: os.Getenv("MINIO_SECRET_KEY"),
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
