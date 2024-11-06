package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	CORS CORS `yaml:"cors"`
	ExternalServices
	Server Server
}

type ExternalServices struct {
	AuthServiceUrl  string
	UsersServiceUrl string
}

type Server struct {
	LogLevel string
	Port     string
	Host     string
}

type CORS struct {
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

func New() *Config {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("can't read config.yml")
	}
	defer f.Close()
	var config Config
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		f.Close()
		log.Fatal("can't decode config.yml") //nolint
	}

	if os.Getenv("LOAD_DOT_ENV") != "false" {
		slog.Info("Loading .env file")
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	es := ExternalServices{
		AuthServiceUrl:  os.Getenv("AUTH_SERVER_URL"),
		UsersServiceUrl: os.Getenv("USERS_SERVER_URL"),
	}
	s := Server{
		LogLevel: os.Getenv("SERVER_LOG_LEVEL"),
		Port:     os.Getenv("SERVER_PORT"),
		Host:     os.Getenv("SERVER_HOST"),
	}
	config.ExternalServices = es
	config.Server = s

	return &config
}
