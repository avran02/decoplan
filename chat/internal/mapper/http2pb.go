package mapper

import (
	"github.com/avran02/decoplan/chat/internal/dto"
	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SaveMessageHttpRequest2Pb(req dto.NewMessageDto, id uint64) *storagepb.Message {
	attachments := make([]*storagepb.Attachment, 0, len(req.Content.Attachments))
	for _, attachment := range req.Content.Attachments {
		attachments = append(attachments, &storagepb.Attachment{
			Id:        attachment.ID,
			Url:       attachment.URL,
			ChatId:    req.ChatID,
			MessageId: id,
		})
	}

	return &storagepb.Message{
		Id:          id,
		ChatId:      req.ChatID,
		Sender:      req.Sender,
		Content:     req.Content.Text,
		Attachments: attachments,
		CreatedAt:   timestamppb.Now(),
		DeletedAt:   nil,
	}
}
