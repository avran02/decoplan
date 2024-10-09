package mapper

import (
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MessageModelToPB(model models.Message) *pb.Message {
	attachments := make([]*pb.Attachment, 0, len(model.Attachments))
	for _, a := range model.Attachments {
		attachments = append(attachments, AttachmentModelToPB(a))
	}
	resp := &pb.Message{
		Id:          &model.ID,
		ChatId:      model.ChatID,
		Sender:      model.Sender,
		Content:     model.Content,
		Attachments: attachments,
		CreatedAt:   timestamppb.New(model.CreatedAt),
	}

	if model.DeletedAt != nil {
		resp.DeletedAt = timestamppb.New(*model.DeletedAt)
	}
	return resp
}

func AttachmentModelToPB(model models.Attachment) *pb.Attachment {
	return &pb.Attachment{
		Id:        model.ID,
		Url:       model.URL,
		ChatId:    model.ChatID,
		MessageId: model.MessageID,
	}
}
