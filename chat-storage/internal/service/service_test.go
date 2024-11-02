package service_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/internal/service"
	"github.com/avran02/decoplan/chat-storage/tests/utils"
	"github.com/stretchr/testify/require"
)

func TestChatService(t *testing.T) {
	ctx := context.Background()
	slog.SetDefault(slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelDebug},
	)))

	// Setup repositories and cleanup
	mongoRepo, mongoTeardown := utils.SetupMongoContainer(t)
	defer mongoTeardown()

	redisRepo, redisTeardown := utils.SetupRedisContainer(t)
	defer redisTeardown()

	// Create the service instance
	srv := service.New(mongoRepo, redisRepo)

	// Define test parameters
	chatID := "chat123"
	message1 := models.Message{ID: 1, ChatID: chatID, Content: "Hello, world!"}
	message2 := models.Message{ID: 2, ChatID: chatID, Content: "How are you?"}

	// Test creating a chat
	err := srv.CreateChat(ctx, chatID)
	require.NoError(t, err)

	// Test saving messages
	err = srv.SaveMessage(ctx, message1)
	require.NoError(t, err)

	err = srv.SaveMessage(ctx, message2)
	require.NoError(t, err)

	// Test retrieving messages
	messages, err := srv.GetMessages(ctx, chatID, 10, 0)
	require.NoError(t, err)
	require.Len(t, messages, 2)
	require.Equal(t, message1.Content, messages[0].Content)
	require.Equal(t, message2.Content, messages[1].Content)

	// Test deleting a message
	err = srv.DeleteMessage(ctx, chatID, message1.ID)
	require.NoError(t, err)

	slog.Info("deleted message", "messageID", message1.ID)
	// Verify the message has been deleted
	messages, err = srv.GetMessages(ctx, chatID, 10, 0)
	require.NoError(t, err)
	require.Len(t, messages, 1)
	require.Equal(t, message2.Content, messages[0].Content)

	// Test caching last messages
	cachedMessages, err := srv.CacheLastMessages(ctx, chatID, 10, 0)
	require.NoError(t, err)
	require.Len(t, cachedMessages, 1)
	require.Equal(t, message2.Content, cachedMessages[0].Content)
}
