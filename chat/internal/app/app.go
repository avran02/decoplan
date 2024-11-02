package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/chat/internal/config"
	"github.com/avran02/decoplan/chat/internal/hub"
	"github.com/avran02/decoplan/chat/internal/router"
	"github.com/avran02/decoplan/chat/internal/service"
	"github.com/avran02/decoplan/chat/logger"
)

type App struct {
	config *config.Config
	router *router.Router
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
	config := config.New()
	logger.Setup(config.Server)
	slog.Debug(fmt.Sprintf("config: %+v", config))

	s := service.New(connectExternalServices(
		config.ExternalServices.AuthURL,
		config.ExternalServices.UsersURL,
		config.ExternalServices.StorageURL,
	))
	hub := hub.New(s)

	router := router.New(hub)
	return &App{
		config: config,
		router: router,
	}
}
