package controller

import (
	"context"
	"log/slog"

	"github.com/avran02/decoplan/chat-storage/internal/mapper"
	"github.com/avran02/decoplan/chat-storage/internal/service"
	"github.com/avran02/decoplan/chat-storage/pb"
)

type Controller interface {
	SaveMessage(ctx context.Context, req *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error)
	GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error)
	DeleteMessage(ctx context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error)
	CacheLastMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error)
	CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error)
}

type controller struct {
	service service.Service
}

func (c *controller) SaveMessage(ctx context.Context, req *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	if err := c.service.SaveMessage(ctx, mapper.FromSaveMessageDtoToModel(req)); err != nil {
		slog.Error("failed to save message", "error", err.Error())
		return nil, err
	}

	return &pb.SaveMessageResponse{
		Ok: true,
	}, nil
}

func (c *controller) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	messages, err := c.service.GetMessages(
		ctx,
		req.ChatId,
		req.Limit,
		req.Offset,
	)
	if err != nil {
		slog.Error("failed to get messages", "error", err.Error())
		return nil, err
	}

	return mapper.FromModelToGetMessagesResponse(messages), nil
}

func (c *controller) DeleteMessage(ctx context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	if err := c.service.DeleteMessage(
		ctx,
		req.ChatId,
		req.MessageId,
	); err != nil {
		slog.Error("failed to delete message", "error", err.Error())
		return nil, err
	}

	return &pb.DeleteMessageResponse{
		Ok: true,
	}, nil
}

func (c *controller) CacheLastMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	messages, err := c.service.CacheLastMessages(
		ctx,
		req.ChatId,
		req.Limit,
		req.Offset,
	)
	if err != nil {
		slog.Error("failed to cache last messages", "error", err.Error())
		return nil, err
	}

	return mapper.FromModelToGetMessagesResponse(messages), nil
}

func (c *controller) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	if err := c.service.CreateChat(ctx, req.ChatId); err != nil {
		slog.Error("failed to create chat", "error", err.Error())
		return nil, err
	}

	return &pb.CreateChatResponse{
		Ok: true,
	}, nil
}

func New(service service.Service) Controller {
	slog.Info("initializing controller")
	return &controller{
		service: service,
	}
}
