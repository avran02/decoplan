package chat_storage_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/avran02/decoplan/chat-storage/pb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test2ChatStorageService(t *testing.T) {
	ts, teardown := setup()
	defer teardown()

	chatID := "test-2-" + uuid.NewString()

	// Test CreateChat
	createChatResp, err := ts.client.CreateChat(context.Background(), &pb.CreateChatRequest{ChatId: chatID})
	assert.NoError(t, err)
	assert.True(t, createChatResp.Ok)

	// Test SaveMessage with valid chat
	message := &pb.Message{
		Id:        1,
		ChatId:    chatID,
		Sender:    "user1",
		Content:   "Hello, World!",
		CreatedAt: timestamppb.New(time.Now()),
	}
	saveMessageResp, err := ts.client.SaveMessage(context.Background(), &pb.SaveMessageRequest{Message: message})
	assert.NoError(t, err)
	assert.True(t, saveMessageResp.Ok)

	// Test SaveMessage with Attachments
	messageWithAttachment := &pb.Message{
		Id:      2,
		ChatId:  chatID,
		Sender:  "user2",
		Content: "Check this out",
		Attachments: []*pb.Attachment{
			{Id: "1", MessageId: 2, ChatId: chatID, Url: "http://example.com/attachment.jpg"},
		},
		CreatedAt: timestamppb.New(time.Now()),
	}
	saveMessageResp, err = ts.client.SaveMessage(context.Background(), &pb.SaveMessageRequest{Message: messageWithAttachment})
	assert.NoError(t, err)
	assert.True(t, saveMessageResp.Ok)

	// Test GetMessages with offset and limit
	getMessagesResp, err := ts.client.GetMessages(context.Background(), &pb.GetMessagesRequest{ChatId: chatID, Limit: 1, Offset: 2})
	assert.NoError(t, err)
	assert.Len(t, getMessagesResp.Messages, 1)
	assert.Equal(t, messageWithAttachment.Content, getMessagesResp.Messages[0].Content)

	// Test Deleting Non-Existent Message
	deleteMessageResp, err := ts.client.DeleteMessage(context.Background(), &pb.DeleteMessageRequest{MessageId: 999, ChatId: chatID})
	assert.NoError(t, err)
	assert.True(t, deleteMessageResp.Ok)

	// Test Duplicate Message Save
	_, err = ts.client.SaveMessage(context.Background(), &pb.SaveMessageRequest{Message: message})
	assert.NoError(t, err)

	// Test CacheLastMessages
	cacheLastMessagesResp, err := ts.client.CacheLastMessages(context.Background(), &pb.GetMessagesRequest{ChatId: chatID, Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Len(t, cacheLastMessagesResp.Messages, 2)

	// Test DeleteMessage
	deleteMessageResp, err = ts.client.DeleteMessage(context.Background(), &pb.DeleteMessageRequest{MessageId: 1, ChatId: chatID})
	assert.NoError(t, err)
	assert.True(t, deleteMessageResp.Ok)

	// Check that the message is deleted
	getMessagesResp, err = ts.client.GetMessages(context.Background(), &pb.GetMessagesRequest{ChatId: chatID, Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Len(t, getMessagesResp.Messages, 1) // Only one message should remain
}
