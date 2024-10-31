package mapper

import (
	"github.com/avran02/decoplan/chat/internal/models"
	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
)

func PbMsgToModel(msg *storagepb.Message) models.Message {
	attachments := make([]models.Attachment, 0, len(msg.Attachments))
	for _, attachment := range msg.Attachments {
		attachments = append(attachments, models.Attachment{
			ID:        attachment.GetId(),
			MessageID: attachment.GetMessageId(),
			URL:       attachment.GetUrl(),
			ChatID:    attachment.GetChatId(),
		})
	}

	deletedAt := msg.GetDeletedAt().AsTime()
	return models.Message{
		ID:     msg.GetId(),
		ChatID: msg.GetChatId(),
		Sender: msg.GetContent(),
		Content: models.Content{
			Attachments: attachments,
			Text:        msg.GetContent(),
		},
		CreatedAt: msg.CreatedAt.AsTime(),
		DeletedAt: &deletedAt,
	}
}
