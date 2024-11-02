package mapper

import (
	"github.com/avran02/decoplan/chat/internal/dto"
	storagepb "github.com/avran02/decoplan/chat/pb/chat_storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SaveMessageHttpRequestToPb(req dto.FromClientMessageDto, sender string) *storagepb.MessageReq {
	attachments := make([]*storagepb.AttachmentReq, 0, len(req.Content.Attachments))
	for _, attachment := range req.Content.Attachments {
		attachments = append(attachments, &storagepb.AttachmentReq{
			Id:  attachment.ID,
			Url: attachment.URL,
		})
	}

	return &storagepb.MessageReq{
		ChatId:      req.ChatID,
		Sender:      sender,
		Content:     req.Content.Text,
		Attachments: attachments,
		CreatedAt:   timestamppb.Now(),
	}
}
