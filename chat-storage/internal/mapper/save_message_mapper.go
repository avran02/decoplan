package mapper

import (
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/pb"
)

func FromSaveMessageDtoToModel(req *pb.SaveMessageRequest) models.Message {
	attachments := make([]models.Attachment, 0, len(req.Message.Attachments))
	for _, a := range req.Message.Attachments {
		attachments = append(attachments, models.Attachment{
			ID:        a.Id,
			URL:       a.Url,
			ChatID:    req.Message.ChatId,
			MessageID: req.Message.Id,
		})
	}

	return models.Message{
		ID:          req.Message.Id,
		ChatID:      req.Message.ChatId,
		Sender:      req.Message.Sender,
		Content:     req.Message.Content,
		CreatedAt:   req.Message.CreatedAt.AsTime(),
		Attachments: attachments,
	}
}
