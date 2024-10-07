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

func setupMongoContainer(t *testing.T) (repository.MongoRepository, func()) {
	// Create a new pool to manage Docker resources
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	// Run a MongoDB container
	resource, err := pool.Run("mongo", "latest", []string{
		"MONGO_INITDB_DATABASE=testdb",
		"MONGO_INITDB_ROOT_USERNAME=123",
		"MONGO_INITDB_ROOT_PASSWORD=123",
	})
	require.NoError(t, err)

	// Connect to MongoDB

	// Create MongoRepository
	repo := repository.NewMongoRepository(&config.Mongo{
		Host:     "localhost",
		Port:     resource.GetPort("27017/tcp"),
		User:     "123",
		Password: "123",
		Database: "testdb",
	})

	// Return the repository and a cleanup function
	return repo, func() {
		repo.Close()
		pool.Purge(resource)
	}
}

func TestMongoRepository_E2E(t *testing.T) {
	ctx := context.Background()

	// Start MongoDB container and get MongoRepository
	repo, tearDown := setupMongoContainer(t)
	defer tearDown()

	// Test 1: Create a chat
	err := repo.CreateChat(ctx, "chat123")
	require.NoError(t, err)

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
