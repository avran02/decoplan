package mapper

import (
	"github.com/avran02/decoplan/gateway/internal/dto"
	"github.com/avran02/decoplan/gateway/pb"
)

func ChatInfoFromPbToHttp(in *pb.GetChatResponse) dto.GetChatResponse {
	members := make([]dto.UserMember, len(in.Members))
	for _, m := range in.GetMembers() {
		members = append(members, dto.UserMember{
			UserID: m.GetUserID(),
		})
	}

	return dto.GetChatResponse{
		ID:       in.Id,
		ChatName: in.ChatName,
		Avatar:   in.Avatar,
		Members:  members,
	}
}
