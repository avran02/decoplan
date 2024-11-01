package app

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/files/internal/config"
	"github.com/avran02/decoplan/files/internal/controller"
	"github.com/avran02/decoplan/files/internal/middleware"
	"github.com/avran02/decoplan/files/internal/router"
	"github.com/avran02/decoplan/files/internal/service"
	"github.com/avran02/decoplan/files/logger"
	"github.com/avran02/decoplan/files/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	authMiddleware := middleware.NewAuthMiddleware(mustConnectExternalServices(conf))
	router := router.New(controller, authMiddleware)

	return &App{
		config: conf,
		router: router,
	}
}

func mustConnectExternalServices(config *config.Config) pb.AuthServiceClient {
	conn, err := grpc.NewClient(config.ExternalServices.AuthServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err)
	}
	return pb.NewAuthServiceClient(conn)
}
