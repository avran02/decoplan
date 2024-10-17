package app

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/decplan/gateway/internal/config"
	"github.com/avran02/decplan/gateway/internal/controllers"
	"github.com/avran02/decplan/gateway/internal/router"
	"github.com/avran02/decplan/gateway/internal/services"
	"github.com/avran02/decplan/gateway/logger"
	"github.com/avran02/decplan/gateway/pb"
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
	service := services.NewAuthService(connectAuthService(conf.AuthServiceUrl))
	controller := controllers.NewAuthController(service)
	router := router.New(controller)

	return &App{
		config: conf,
		router: router,
	}
}

func connectAuthService(endpoint string) pb.AuthServiceClient {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err)
	}
	return pb.NewAuthServiceClient(conn)
}
