package service

import (
	"context"
	"fmt"

	"github.com/avran02/decoplan/chat/internal/models"
	authpb "github.com/avran02/decoplan/chat/pb/auth"
	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
	userspb "github.com/avran02/decoplan/chat/pb/users"
)

type Service interface {
	SaveMessage(ctx context.Context, message *storagepb.Message) error
	DeleteMessage(ctx context.Context, chatID string, messageID uint64) error
	GetMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error)
	GetChatMembers(ctx context.Context, chatID, userID string) ([]string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

type service struct {
	storageClient storagepb.ChatStorageServiceClient
	usersClient   userspb.UsersServiceClient
	authClient    authpb.AuthServiceClient
}

func (s *service) SaveMessage(ctx context.Context, message *storagepb.Message) error {
	resp, err := s.storageClient.SaveMessage(ctx, &storagepb.SaveMessageRequest{Message: message})
	if err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}
	if resp.GetOk() == false {
		return ErrUnknownError
	}

	return nil
}

func (s *service) DeleteMessage(ctx context.Context, chatID string, messageID uint64) error {
	resp, err := s.storageClient.DeleteMessage(ctx, &storagepb.DeleteMessageRequest{
		ChatId:    chatID,
		MessageId: messageID},
	)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	if resp.GetOk() == false {
		return ErrUnknownError
	}

	return nil
}

func (s *service) GetMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error) {
	resp, err := s.storageClient.GetMessages(ctx, &storagepb.GetMessagesRequest{
		ChatId: chatID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	messages := make([]models.Message, 0, len(resp.GetMessages()))
	for _, message := range resp.GetMessages() {
		messages = append(messages, models.Message{
			ID:        message.GetId(),
			ChatID:    message.GetChatId(),
			Content:   message.GetContent(),
			CreatedAt: message.GetCreatedAt().AsTime(),
		})
	}
	return messages, nil
}

func (s *service) GetChatMembers(ctx context.Context, chatID, userID string) ([]string, error) {
	resp, err := s.usersClient.GetChat(ctx, &userspb.GetChatRequest{Id: chatID})
	if err != nil {
		return nil, fmt.Errorf("failed to get chat members: %w", err)
	}
	userFound := false
	members := make([]string, 0, len(resp.GetMembers()))
	for _, member := range resp.GetMembers() {
		if member.GetUserID() == userID {
			userFound = true
		}
		members = append(members, member.GetUserID())
	}
	if !userFound {
		return nil, ErrChatForbidden
	}
	return members, nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	resp, err := s.authClient.ValidateToken(context.Background(), &authpb.ValidateTokenRequest{AccessToken: token})
	if err != nil {
		return "", fmt.Errorf("failed to validate token: %w", err)
	}
	return resp.GetId(), nil
}

func New(storageClient storagepb.ChatStorageServiceClient) Service {
	return &service{
		storageClient: storageClient,
	}
}
