package mapper

import (
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// WARNING! does not map ID
func FromSaveMessageDtoToModel(req *pb.SaveMessageRequest, msgID uint64) models.Message {
	attachments := make([]models.Attachment, 0, len(req.Message.Attachments))
	for _, a := range req.Message.Attachments {
		attachments = append(attachments, models.Attachment{
			MessageID: msgID,
			ID:        a.Id,
			URL:       a.Url,
			ChatID:    req.Message.ChatId,
		})
	}

	return models.Message{
		ID:          msgID,
		ChatID:      req.Message.ChatId,
		Sender:      req.Message.Sender,
		Content:     req.Message.Content,
		CreatedAt:   req.Message.CreatedAt.AsTime(),
		Attachments: attachments,
	}
}

func FromModelToSaveMessageResponse(model models.Message) *pb.SaveMessageResponse {
	attachments := make([]*pb.Attachment, 0, len(model.Attachments))
	for _, a := range model.Attachments {
		attachments = append(attachments, AttachmentModelToPB(a))
	}
	return &pb.SaveMessageResponse{
		Message: &pb.Message{
			Id:          model.ID,
			ChatId:      model.ChatID,
			Sender:      model.Sender,
			Content:     model.Content,
			Attachments: attachments,
			CreatedAt:   timestamppb.New(model.CreatedAt),
			DeletedAt:   nil,
		},
	}
}
