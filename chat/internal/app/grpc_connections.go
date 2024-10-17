package app

import (
	"log"

	authpb "github.com/avran02/decoplan/chat/pb/auth"
	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
	userspb "github.com/avran02/decoplan/chat/pb/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectExternalServices(authEndpoint, usersEndpoint, storageEndpoint string) (storagepb.ChatStorageServiceClient, authpb.AuthServiceClient, userspb.UsersServiceClient) {
	authConn, err := grpc.NewClient(authEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to auth service: %s", err)
	}

	usersConn, err := grpc.NewClient(usersEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to users service: %s", err)
	}

	storageConn, err := grpc.NewClient(storageEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to chat storage service: %s", err)
	}

	return storagepb.NewChatStorageServiceClient(storageConn), authpb.NewAuthServiceClient(authConn), userspb.NewUsersServiceClient(usersConn)
}
