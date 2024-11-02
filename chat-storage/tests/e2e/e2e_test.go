package chat_storage_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/avran02/decoplan/chat-storage/pb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestSuite holds the necessary data for our e2e tests.
type TestSuite struct {
	conn   *grpc.ClientConn
	client pb.ChatStorageServiceClient
}

func setup() (*TestSuite, func()) {
	// Set up a connection to the gRPC server
	conn, err := grpc.NewClient("localhost:51051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewChatStorageServiceClient(conn)

	// Return the test suite and teardown function
	return &TestSuite{conn: conn, client: client}, func() {
		conn.Close()
	}
}

func TestChatStorageService(t *testing.T) {
	ts, teardown := setup()
	defer teardown()

	// Test CreateChat
	chatID := "test-" + uuid.NewString()
	createChatResp, err := ts.client.CreateChat(context.Background(), &pb.CreateChatRequest{ChatId: chatID})
	assert.NoError(t, err)
	assert.True(t, createChatResp.Ok)

	// Test SaveMessage
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

	// Test GetMessages
	getMessagesResp, err := ts.client.GetMessages(context.Background(), &pb.GetMessagesRequest{ChatId: chatID, Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.NotEmpty(t, getMessagesResp.Messages)

	// Test DeleteMessage
	deleteMessageResp, err := ts.client.DeleteMessage(context.Background(), &pb.DeleteMessageRequest{MessageId: 1, ChatId: chatID})
	assert.NoError(t, err)
	assert.True(t, deleteMessageResp.Ok)

	// Check that the message is deleted
	getMessagesResp, err = ts.client.GetMessages(context.Background(), &pb.GetMessagesRequest{ChatId: chatID, Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Empty(t, getMessagesResp.Messages)
}
