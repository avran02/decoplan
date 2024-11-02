package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/logger"
	"github.com/avran02/decoplan/chat-storage/tests/utils"
	"github.com/stretchr/testify/require"
)

func TestMongoRepository_E2E(t *testing.T) {
	ctx := context.Background()
	logger.Setup(config.Server{LogLevel: "debug"})

	// Start MongoDB container and get MongoRepository
	repo, tearDown := utils.SetupMongoContainer(t)
	defer tearDown()

	// Test 1: Create a chat
	err := repo.CreateChat(ctx, "chat123")
	require.NoError(t, err)

	// Create sample messages
	message1 := models.Message{
		ID:      0,
		ChatID:  "chat123",
		Sender:  "user1",
		Content: "First message",
		Attachments: []models.Attachment{
			{
				ID:        "1",
				URL:       "https://example.com/image.jpg",
				ChatID:    "chat123",
				MessageID: 0,
			},
		},
		CreatedAt: time.Now(),
	}
	message2 := models.Message{
		ID:        1,
		ChatID:    "chat123",
		Sender:    "user2",
		Content:   "Second message",
		CreatedAt: time.Now(),
	}

	// Test 2: Save messages
	err = repo.SaveMessage(ctx, message1)
	require.NoError(t, err)
	err = repo.SaveMessage(ctx, message2)
	require.NoError(t, err)

	// Test 3: Validate saved messages
	savedMessages, err := repo.GetMessages(ctx, "chat123", 0, 1)
	require.NoError(t, err)
	require.Len(t, savedMessages, 2)
	require.Equal(t, message1.Content, savedMessages[0].Content)
	require.Equal(t, message2.Content, savedMessages[1].Content)
	require.Equal(t, "https://example.com/image.jpg", savedMessages[0].Attachments[0].URL)

	// Test 4: Delete a message
	err = repo.DeleteMessage(ctx, "chat123", 1)
	require.NoError(t, err)

	// Test 5: Validate message deletion
	savedMessagesAfterDeletion, err := repo.GetMessages(ctx, "chat123", 0, 1)
	require.NoError(t, err)
	require.Len(t, savedMessagesAfterDeletion, 1)
	require.Equal(t, message1.Content, savedMessagesAfterDeletion[0].Content)

	// Clean up and close MongoDB connection
	err = repo.Close()
	require.NoError(t, err)
}
