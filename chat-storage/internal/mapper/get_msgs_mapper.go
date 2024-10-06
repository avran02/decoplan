package mapper

import (
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/pb"
)

func FromModelToGetMessagesResponse(model []models.Message) *pb.GetMessagesResponse {
	resp := &pb.GetMessagesResponse{
		Messages: make([]*pb.Message, 0, len(model)),
	}

	for _, m := range model {
		resp.Messages = append(resp.Messages, MessageModelToPB(m))
	}
	return resp
}
