package tests

import (
	"context"
	"testing"
	"time"

	"github.com/avran02/decoplan/users/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserScenario(t *testing.T) {
	client := setupClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Step 1: Create two users
	user1 := &pb.CreateUserRequest{
		Id:        "user-1-id",
		Name:      "Alice",
		BirthDate: timestamppb.Now(),
	}
	user2 := &pb.CreateUserRequest{
		Id:        "user-2-id",
		Name:      "Bob",
		BirthDate: timestamppb.Now(),
	}

	user1Resp, err := client.CreateUser(ctx, user1)
	assert.NoError(t, err)
	assert.True(t, user1Resp.Ok)

	user2Resp, err := client.CreateUser(ctx, user2)
	assert.NoError(t, err)
	assert.True(t, user2Resp.Ok)

	// Step 2: Create a chat with user1 as an initial member
	chatReq := &pb.CreateChatRequest{
		Name:    "Friends Chat",
		UserIDs: []string{"user-1-id"},
	}
	chatResp, err := client.CreateChat(ctx, chatReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, chatResp.ChatID)

	chatID := chatResp.ChatID

	// Step 3: Add user2 to the chat
	addUserReq := &pb.AddUserToChatRequest{
		ChatID: chatID,
		UserID: "user-2-id",
	}
	addUserResp, err := client.AddUserToChat(ctx, addUserReq)
	assert.NoError(t, err)
	assert.True(t, addUserResp.Ok)

	// Step 4: Get the chat details to verify both users are present
	getChatReq := &pb.GetChatRequest{Id: chatID}
	getChatResp, err := client.GetChat(ctx, getChatReq)
	assert.NoError(t, err)
	chatName := getChatResp.ChatName
	assert.Equal(t, "Friends Chat", *chatName)
	assert.Len(t, getChatResp.Members, 2)

	// Step 5: Remove user2 from the chat
	removeUserReq := &pb.RemoveUserFromChatRequest{
		ChatID: chatID,
		UserID: "user-2-id",
	}
	removeUserResp, err := client.RemoveUserFromChat(ctx, removeUserReq)
	assert.NoError(t, err)
	assert.True(t, removeUserResp.Ok)

	// Step 6: Get the chat details again to verify that user2 is removed
	getChatRespAfterRemoval, err := client.GetChat(ctx, getChatReq)
	assert.NoError(t, err)
	assert.Len(t, getChatRespAfterRemoval.Members, 1)
	assert.Equal(t, "user-1-id", getChatRespAfterRemoval.Members[0].UserID)

	// Step 7: Delete the chat
	deleteChatReq := &pb.DeleteChatRequest{Id: chatID}
	deleteChatResp, err := client.DeleteChat(ctx, deleteChatReq)
	assert.NoError(t, err)
	assert.True(t, deleteChatResp.Ok)

	// Step 8: Verify the chat has been deleted by trying to get it again
	getChatRespAfterDeletion, err := client.GetChat(ctx, getChatReq)
	assert.Error(t, err)
	assert.Nil(t, getChatRespAfterDeletion)
}
