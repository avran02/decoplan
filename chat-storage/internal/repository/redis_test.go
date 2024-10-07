package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/internal/repository"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

func setupRedisContainer(t *testing.T) (repository.RedisRepository, func()) {
	// Create a new pool to manage Docker resources
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	err = pool.Client.Ping()
	require.NoError(t, err)

	// Run a Redis container
	resource, err := pool.Run("redis", "latest", []string{})
	require.NoError(t, err)

	client := repository.NewRedisRepository(&config.Redis{
		Host:     "localhost",
		Port:     resource.GetPort("6379/tcp"),
		Password: "",
		Database: 0,
	})

	// Return the client and cleanup function
	return client, func() {
		pool.Purge(resource) // Clean up container after test
	}
}

func TestRedisRepository_E2E(t *testing.T) {
	ctx := context.Background()

	// Start Redis container and get RedisRepository
	repo, tearDown := setupRedisContainer(t)
	defer tearDown()

	// Create sample messages
	message1 := models.Message{
		ID:        0,
		ChatID:    "chat123",
		Sender:    "user1",
		Content:   "First message",
		CreatedAt: time.Now(),
	}
	message2 := models.Message{
		ID:        1,
		ChatID:    "chat123",
		Sender:    "user2",
		Content:   "Second message",
		CreatedAt: time.Now(),
	}

	// Test 1: Cache group of last messages
	err := repo.CacheGroupLastMessages(ctx, "chat123", []models.Message{message1, message2})
	require.NoError(t, err)

	// Test 2: Validate saved messages
	savedMessages, err := repo.GetMessages(ctx, "chat123", 0, 1)
	require.NoError(t, err)
	require.Len(t, savedMessages, 2)
	require.Equal(t, message1.Content, savedMessages[0].Content)
	require.Equal(t, message2.Content, savedMessages[1].Content)

	// Test 3: Validate cache limits
	startIdx, endIdx, err := repo.GetCacheLimits(ctx, "chat123")
	require.NoError(t, err)
	require.Equal(t, uint64(1), startIdx)
	require.Equal(t, uint64(2), endIdx)

	// Test 4: Delete a message
	err = repo.DeleteMessage(ctx, "chat123", 0)
	require.NoError(t, err)

	// Test 5: Validate message deletion
	savedMessagesAfterDeletion, err := repo.GetMessages(ctx, "chat123", 0, 1)
	require.NoError(t, err)
	require.Len(t, savedMessagesAfterDeletion, 1)
	require.Equal(t, message2.Content, savedMessagesAfterDeletion[0].Content)

	// Clean up and close Redis
	err = repo.Close()
	require.NoError(t, err)
}
