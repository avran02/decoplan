package app

import (
	"log/slog"
	"net"
	"os"

	"github.com/avran02/decoplan/users/internal/config"
	"github.com/avran02/decoplan/users/internal/controller"
	"github.com/avran02/decoplan/users/internal/repository"
	"github.com/avran02/decoplan/users/internal/server"
	"github.com/avran02/decoplan/users/internal/service"
	"github.com/avran02/decoplan/users/logger"
	"github.com/avran02/decoplan/users/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var opts []grpc.ServerOption

type App struct {
	Config *config.Config
	Server server.UsersServer
}

func (app *App) Run() {
	host := app.Config.Server.Host + ":" + app.Config.Server.Port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		slog.Error("failed to listen:\n" + err.Error())
		os.Exit(1)
	}

	slog.Info("Listening on " + host)

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersServiceServer(grpcServer, app.Server)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("usersservice", grpc_health_v1.HealthCheckResponse_SERVING)

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("failed to serve:\n" + err.Error())
		os.Exit(1)
	}
}

func New() *App {
	conf := config.New()
	logger.Setup(conf.Server)
	repository := repository.New(conf.DB)
	service := service.New(repository)
	controller := controller.New(service)
	server := server.New(controller)

	return &App{
		Config: conf,
		Server: server,
	}
}
