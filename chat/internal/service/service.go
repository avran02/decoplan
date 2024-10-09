package service

import (
	"context"
	"fmt"

	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
	userspb "github.com/avran02/decoplan/chat/pb/users"
)

type Service interface {
	SaveMessage(ctx context.Context, message storagepb.Message) error
}

type service struct {
	storageClient storagepb.ChatStorageServiceClient
	usersClient   userspb.UsersServiceClient
}

func (s *service) SaveMessage(ctx context.Context, message storagepb.Message) error {
	resp, err := s.storageClient.SaveMessage(ctx, &storagepb.SaveMessageRequest{Message: &message})
	if err != nil || resp.GetOk() == false {
		return fmt.Errorf("failed to save message: %w", err)
	}

	return nil
}

func New(storageClient storagepb.ChatStorageServiceClient) Service {
	return &service{
		storageClient: storageClient,
	}
}
