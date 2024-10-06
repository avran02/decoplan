package server

import (
	"context"
	"log/slog"

	"github.com/avran02/decoplan/chat-storage/internal/controller"
	"github.com/avran02/decoplan/chat-storage/pb"
)

type Server struct {
	pb.UnimplementedChatStorageServiceServer
	Controller controller.Controller
}

func (s Server) SaveMessage(ctx context.Context, req *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	return s.Controller.SaveMessage(ctx, req)
}

func (s Server) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return s.Controller.GetMessages(ctx, req)
}

func (s Server) DeleteMessage(ctx context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	return s.Controller.DeleteMessage(ctx, req)
}

func (s Server) CacheLastMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return s.Controller.CacheLastMessages(ctx, req)
}

func (s Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	return s.Controller.CreateChat(ctx, req)
}

func New(controller controller.Controller) Server {
	slog.Info("initializing server")
	return Server{
		Controller: controller,
	}
}
