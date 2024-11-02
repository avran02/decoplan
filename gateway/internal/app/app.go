package app

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/decoplan/gateway/internal/config"
	"github.com/avran02/decoplan/gateway/internal/controllers"
	"github.com/avran02/decoplan/gateway/internal/router"
	"github.com/avran02/decoplan/gateway/internal/services"
	"github.com/avran02/decoplan/gateway/logger"
	"github.com/avran02/decoplan/gateway/pb"
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
	srv := services.NewUsersService(connectUsersService(conf.ExternalServices.UsersServiceUrl))
	controller := controllers.New(srv)
	router := router.New(controller, connectAuthService(conf.ExternalServices.AuthServiceUrl))

	return &App{
		config: conf,
		router: router,
	}
}

func connectUsersService(endpoint string) pb.UsersServiceClient {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to users service: %s", err)
	}

	return pb.NewUsersServiceClient(conn)
}

func connectAuthService(endpoint string) pb.AuthServiceClient {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err)
	}

	return pb.NewAuthServiceClient(conn)
}
