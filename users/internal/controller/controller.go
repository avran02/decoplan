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

func (c *UserController) AddUserToGroup(ctx context.Context, req *pb.AddUserToGroupRequest) (*pb.AddUserToGroupResponse, error) {
	if err := c.service.AddUserToGroup(ctx, models.UserGroup{
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}); err != nil {
		return nil, err
	}

	return &pb.AddUserToGroupResponse{Ok: true}, nil
}

func (c *UserController) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	groupID, err := c.service.CreateGroup(ctx, req.GetName(), req.GetUserIDs())
	if err != nil {
		return nil, err
	}

	return &pb.CreateGroupResponse{GroupID: groupID}, nil
}

func (c *UserController) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	if err := c.service.DeleteGroup(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &pb.DeleteGroupResponse{Ok: true}, nil
}

func (c *UserController) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	group, err := c.service.GetGroup(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	members := make([]*pb.UserMember, 0)
	for _, v := range group.Members {
		members = append(members, &pb.UserMember{
			UserID: v.ID,
		})
	}

	return &pb.GetGroupResponse{
		Id:        group.ID,
		GroupName: &group.Name,
		Avatar:    group.Avatar,
		Members:   members,
	}, nil
}

func (c *UserController) RemoveUserFromGroup(ctx context.Context, req *pb.RemoveUserFromGroupRequest) (*pb.RemoveUserFromGroupResponse, error) {
	if err := c.service.RemoveUserFromGroup(ctx, models.UserGroup{
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}); err != nil {
		return nil, err
	}

	return &pb.RemoveUserFromGroupResponse{Ok: true}, nil
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
