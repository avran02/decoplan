package server

import (
	"context"

	"github.com/avran02/decoplan/users/internal/controller"

	"github.com/avran02/decoplan/users/pb"
)

type UsersServer struct {
	pb.UnimplementedUsersServiceServer
	*controller.UserController
}

func (s UsersServer) AddUserToGroup(ctx context.Context, req *pb.AddUserToGroupRequest) (*pb.AddUserToGroupResponse, error) {
	return s.UserController.AddUserToGroup(ctx, req)
}

func (s UsersServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	return s.UserController.CreateGroup(ctx, req)
}

func (s UsersServer) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	return s.UserController.DeleteGroup(ctx, req)
}
func (s UsersServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.UserController.CreateUser(ctx, req)
}

func (s UsersServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.UserController.DeleteUser(ctx, req)
}

func (s UsersServer) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	return s.UserController.GetGroup(ctx, req)
}

func (s UsersServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.UserController.GetUser(ctx, req)
}

func (s UsersServer) RemoveUserFromGroup(ctx context.Context, req *pb.RemoveUserFromGroupRequest) (*pb.RemoveUserFromGroupResponse, error) {
	return s.UserController.RemoveUserFromGroup(ctx, req)
}

func (s UsersServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.UserController.UpdateUser(ctx, req)
}

func New(controller *controller.UserController) UsersServer {
	return UsersServer{
		UserController: controller,
	}
}
