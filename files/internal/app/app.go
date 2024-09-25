package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/files/internal/config"
	"github.com/avran02/decoplan/files/internal/controller"
	"github.com/avran02/decoplan/files/internal/router"
	"github.com/avran02/decoplan/files/internal/service"
	"github.com/avran02/decoplan/files/logger"
)

type App struct {
	config *config.Config
	router router.Router
}

func (a *App) Run() error {
	serverEndpoint := fmt.Sprintf("%s:%s", a.config.Server.Host, a.config.Server.Port)
	slog.Info("Starting server at " + serverEndpoint)
	s := http.Server{ //nolint:gosec
		Addr:    serverEndpoint,
		Handler: a.router,
	}

	return s.ListenAndServe()
}

func New() *App {
	conf := config.New()
	logger.Setup(conf.Server)
	service := service.New(conf.Minio)
	controller := controller.New(service)
	router := router.New(controller)

	return &App{
		config: conf,
		router: router,
	}
}
