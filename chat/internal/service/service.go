package service

import "github.com/avran02/decoplan/chat/pb"

type Service interface{}

type service struct {
	storageClient pb.ChatStorageServiceClient
}

func New(storageClient pb.ChatStorageServiceClient) Service {
	return &service{
		storageClient: storageClient,
	}
}
