package controller

import (
	"context"

	"github.com/avran02/decoplan/users/internal/models"
	"github.com/avran02/decoplan/users/internal/service"
	"github.com/avran02/decoplan/users/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserController struct {
	service service.UserService
}

func (c *UserController) AddUserToChat(ctx context.Context, req *pb.AddUserToChatRequest) (*pb.AddUserToChatResponse, error) {
	if err := c.service.AddUserToChat(ctx, models.UserChat{
		ChatID: req.ChatID,
		UserID: req.UserID,
	}); err != nil {
		return nil, err
	}

	return &pb.AddUserToChatResponse{Ok: true}, nil
}

func (c *UserController) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	chatID, err := c.service.CreateChat(ctx, req.GetName(), req.GetUserIDs())
	if err != nil {
		return nil, err
	}

	return &pb.CreateChatResponse{ChatID: chatID}, nil
}

func (c *UserController) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*pb.DeleteChatResponse, error) {
	if err := c.service.DeleteChat(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &pb.DeleteChatResponse{Ok: true}, nil
}

func (c *UserController) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	chat, err := c.service.GetChat(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	members := make([]*pb.UserMember, 0)
	for _, v := range chat.Members {
		members = append(members, &pb.UserMember{
			UserID: v.ID,
		})
	}

	return &pb.GetChatResponse{
		Id:       chat.ID,
		ChatName: &chat.Name,
		Avatar:   chat.Avatar,
		Members:  members,
	}, nil
}

func (c *UserController) RemoveUserFromChat(ctx context.Context, req *pb.RemoveUserFromChatRequest) (*pb.RemoveUserFromChatResponse, error) {
	if err := c.service.RemoveUserFromChat(ctx, models.UserChat{
		ChatID: req.ChatID,
		UserID: req.UserID,
	}); err != nil {
		return nil, err
	}

	return &pb.RemoveUserFromChatResponse{Ok: true}, nil
}
func (c *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := c.service.CreateUser(ctx, req.GetId(), req.GetName(), req.BirthDate.AsTime()); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Ok: true}, nil
}

func (c *UserController) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := c.service.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Avatar:    user.Avatar,
		BirthDate: timestamppb.New(user.BirthDate),
	}, nil
}

func (c *UserController) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	name := req.Name
	birthDate := req.GetBirthDate().AsTime()
	avatar := req.Avatar
	if err := c.service.UpdateUser(ctx, models.UpdateUser{
		ID:        req.GetId(),
		Name:      name,
		BirthDate: &birthDate,
		Avatar:    avatar,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{Ok: true}, nil
}

func (c *UserController) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := c.service.DeleteUser(ctx, req.GetUserID()); err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Ok: true}, nil
}

func New(service service.UserService) *UserController {
	return &UserController{service: service}
}
