package app

import (
	"log/slog"
	"net"
	"os"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/controller"
	"github.com/avran02/decoplan/chat-storage/internal/repository"

	"github.com/avran02/decoplan/chat-storage/internal/server"
	"github.com/avran02/decoplan/chat-storage/internal/service"
	"github.com/avran02/decoplan/chat-storage/logger"
	"github.com/avran02/decoplan/chat-storage/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var opts []grpc.ServerOption

type App struct {
	Config *config.Config
	Server server.Server
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
	pb.RegisterChatStorageServiceServer(grpcServer, app.Server)

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

	slog.Debug("config loaded", "config", conf)

	mongo := repository.NewMongoRepository(&conf.Mongo)
	redis := repository.NewRedisRepository(&conf.Redis)

	service := service.New(mongo, redis)
	controller := controller.New(service)
	server := server.New(controller)

	return &App{
		Config: conf,
		Server: server,
	}
}
