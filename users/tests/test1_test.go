package tests

import (
	"context"
	"testing"
	"time"

	"github.com/avran02/decoplan/users/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const serverAddress = "localhost:50051"

func setupClient(t *testing.T) pb.UsersServiceClient {
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	t.Cleanup(func() {
		conn.Close()
	})
	return pb.NewUsersServiceClient(conn)
}

func TestCreateUser(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CreateUserRequest{
		Id:        "user-1",
		Name:      "John-Doe",
		BirthDate: timestamppb.Now(),
	}

	resp, err := client.CreateUser(ctx, req)
	assert.NoError(t, err)
	assert.True(t, resp.Ok)
}

func TestGetUser(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createReq := &pb.CreateUserRequest{
		Id:        "user-2",
		Name:      "Jane-Doe",
		BirthDate: timestamppb.Now(),
	}
	client.CreateUser(ctx, createReq)

	getReq := &pb.GetUserRequest{Id: "user-2"}
	resp, err := client.GetUser(ctx, getReq)
	assert.NoError(t, err)
	assert.Equal(t, "user-2", resp.Id)
	assert.Equal(t, "Jane-Doe", resp.Name)
	assert.NotNil(t, resp.BirthDate)
}

func TestUpdateUser(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createReq := &pb.CreateUserRequest{
		Id:        "user-3",
		Name:      "Original Name",
		BirthDate: timestamppb.Now(),
	}
	client.CreateUser(ctx, createReq)

	name := "Updated Name"
	updateReq := &pb.UpdateUserRequest{
		Id:   "user-3",
		Name: &name,
	}
	updateResp, err := client.UpdateUser(ctx, updateReq)
	assert.NoError(t, err)
	assert.True(t, updateResp.Ok)

	getReq := &pb.GetUserRequest{Id: "user-3"}
	getResp, err := client.GetUser(ctx, getReq)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", getResp.Name)
}

func TestDeleteUser(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createReq := &pb.CreateUserRequest{
		Id:        "user-4",
		Name:      "To Be Deleted",
		BirthDate: timestamppb.Now(),
	}
	client.CreateUser(ctx, createReq)

	deleteReq := &pb.DeleteUserRequest{UserID: "user-4"}
	deleteResp, err := client.DeleteUser(ctx, deleteReq)
	assert.NoError(t, err)
	assert.True(t, deleteResp.Ok)

	getReq := &pb.GetUserRequest{Id: "user-4"}
	getResp, err := client.GetUser(ctx, getReq)
	assert.Error(t, err)
	assert.Nil(t, getResp)
}

func TestCreateChat(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CreateChatRequest{
		Name:    "New Chat",
		UserIDs: []string{"user-1", "user-2"},
	}

	resp, err := client.CreateChat(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.ChatID)
}

func TestGetChat(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createReq := &pb.CreateChatRequest{
		Name:    "Test Chat",
		UserIDs: []string{"user-1", "user-2"},
	}
	createResp, _ := client.CreateChat(ctx, createReq)

	getReq := &pb.GetChatRequest{Id: createResp.ChatID}
	getResp, err := client.GetChat(ctx, getReq)
	assert.NoError(t, err)
	assert.Equal(t, "Test Chat", *getResp.ChatName)
	assert.Len(t, getResp.Members, 2)
}

func TestAddUserToChat(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createChatReq := &pb.CreateChatRequest{
		Name:    "Chat for Adding User",
		UserIDs: []string{"user-1"},
	}
	createChatResp, _ := client.CreateChat(ctx, createChatReq)

	addUserReq := &pb.AddUserToChatRequest{
		ChatID: createChatResp.ChatID,
		UserID: "user-2",
	}
	addUserResp, err := client.AddUserToChat(ctx, addUserReq)
	assert.NoError(t, err)
	assert.True(t, addUserResp.Ok)
}

func TestRemoveUserFromChat(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createChatReq := &pb.CreateChatRequest{
		Name:    "Chat for Removing User",
		UserIDs: []string{"user-1", "user-2"},
	}
	createChatResp, _ := client.CreateChat(ctx, createChatReq)

	removeUserReq := &pb.RemoveUserFromChatRequest{
		ChatID: createChatResp.ChatID,
		UserID: "user-2",
	}
	removeUserResp, err := client.RemoveUserFromChat(ctx, removeUserReq)
	assert.NoError(t, err)
	assert.True(t, removeUserResp.Ok)
}
