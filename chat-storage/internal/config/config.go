package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server Server
	Mongo  Mongo
	Redis  Redis
}

type Server struct {
	LogLevel string
	Port     string
	Host     string
}

type Mongo struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Database int
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
		Mongo: Mongo{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Database: os.Getenv("MONGO_DATABASE"),
		},
		Redis: Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Database: 0,
		},
	}

	return config
}
