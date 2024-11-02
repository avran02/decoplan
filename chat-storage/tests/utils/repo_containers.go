package utils

import (
	"testing"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/repository"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

func SetupMongoContainer(t *testing.T) (repository.MongoRepository, func()) {
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

func SetupRedisContainer(t *testing.T) (repository.RedisRepository, func()) {
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
