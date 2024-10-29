package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/avran02/decplan/gateway/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UsersService interface {
	CreateChat(ctx context.Context, name string, userIDs []string) (string, error)
	GetChat(ctx context.Context, id string) (*pb.GetChatResponse, error)
	DeleteChat(ctx context.Context, id string) error
	RemoveUserFromChat(ctx context.Context, chatID, userID string) error
	AddUserToChat(ctx context.Context, chatID, userID string) error
	CreateUser(ctx context.Context, id, name string, birthDate time.Time) error
	GetUser(ctx context.Context, id string) (*pb.GetUserResponse, error)
	UpdateUser(ctx context.Context, id string, name, avatar *string, birthDate *time.Time) error
	DeleteUser(ctx context.Context, id string) error
}

type usersService struct {
	client pb.UsersServiceClient
}

func (s *usersService) CreateChat(ctx context.Context, name string, userIDs []string) (string, error) {
	slog.Info("usersService.CreateChat")
	req := &pb.CreateChatRequest{
		Name:    name,
		UserIDs: userIDs,
	}
	resp, err := s.client.CreateChat(ctx, req)
	if err != nil {
		return "", fmt.Errorf("can't make gRPC call: %w", err)
	}

	return resp.ChatID, nil
}

func (s *usersService) GetChat(ctx context.Context, id string) (*pb.GetChatResponse, error) {
	slog.Info("usersService.GetChat")
	req := &pb.GetChatRequest{
		Id: id,
	}
	resp, err := s.client.GetChat(ctx, req)
	if err != nil || resp == nil {
		return nil, fmt.Errorf("can't make gRPC call: %w", err)
	}
	return resp, nil
}

func (s *usersService) DeleteChat(ctx context.Context, id string) error {
	slog.Info("usersService.DeleteChat")
	req := &pb.DeleteChatRequest{
		Id: id,
	}
	resp, err := s.client.DeleteChat(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func (s *usersService) RemoveUserFromChat(ctx context.Context, chatID, userID string) error {
	slog.Info("usersService.RemoveUserFromChat")
	req := &pb.RemoveUserFromChatRequest{
		ChatID: chatID,
		UserID: userID,
	}
	resp, err := s.client.RemoveUserFromChat(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func (s *usersService) AddUserToChat(ctx context.Context, chatID, userID string) error {
	slog.Info("usersService.AddUserToChat")
	req := &pb.AddUserToChatRequest{
		UserID: userID,
		ChatID: chatID,
	}
	resp, err := s.client.AddUserToChat(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func (s *usersService) CreateUser(ctx context.Context, id, name string, birthDate time.Time) error {
	slog.Info("usersService.CreateUser")
	req := &pb.CreateUserRequest{
		Id:        id,
		Name:      name,
		BirthDate: timestamppb.New(birthDate),
	}
	resp, err := s.client.CreateUser(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func (s *usersService) GetUser(ctx context.Context, id string) (*pb.GetUserResponse, error) {
	slog.Info("usersService.GetUser")
	req := &pb.GetUserRequest{
		Id: id,
	}
	resp, err := s.client.GetUser(ctx, req)
	if err != nil || resp == nil {
		return nil, fmt.Errorf("can't make gRPC call: %w", err)
	}
	return resp, nil
}

func (s *usersService) UpdateUser(ctx context.Context, id string, name, avatar *string, birthDate *time.Time) error {
	slog.Info("usersService.UpdateUser")
	req := &pb.UpdateUserRequest{
		Id:        id,
		Name:      name,
		Avatar:    avatar,
		BirthDate: timestamppb.New(*birthDate),
	}
	resp, err := s.client.UpdateUser(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func (s *usersService) DeleteUser(ctx context.Context, id string) error {
	slog.Info("usersService.DeleteUser")
	req := &pb.DeleteUserRequest{
		UserID: id,
	}
	resp, err := s.client.DeleteUser(ctx, req)
	if err != nil || resp == nil || !resp.Ok {
		return fmt.Errorf("can't make gRPC call: %w", err)
	}
	return nil
}

func NewUsersService(client pb.UsersServiceClient) UsersService {
	return &usersService{
		client: client,
	}
}
