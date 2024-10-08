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

func (s UsersServer) AddUserToChat(ctx context.Context, req *pb.AddUserToChatRequest) (*pb.AddUserToChatResponse, error) {
	return s.UserController.AddUserToChat(ctx, req)
}

func (s UsersServer) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	return s.UserController.CreateChat(ctx, req)
}

func (s UsersServer) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*pb.DeleteChatResponse, error) {
	return s.UserController.DeleteChat(ctx, req)
}
func (s UsersServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.UserController.CreateUser(ctx, req)
}

func (s UsersServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.UserController.DeleteUser(ctx, req)
}

func (s UsersServer) GetChat(ctx context.Context, req *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	return s.UserController.GetChat(ctx, req)
}

func (s UsersServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.UserController.GetUser(ctx, req)
}

func (s UsersServer) RemoveUserFromChat(ctx context.Context, req *pb.RemoveUserFromChatRequest) (*pb.RemoveUserFromChatResponse, error) {
	return s.UserController.RemoveUserFromChat(ctx, req)
}

func (s UsersServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.UserController.UpdateUser(ctx, req)
}

func New(controller *controller.UserController) UsersServer {
	return UsersServer{
		UserController: controller,
	}
}
