package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository interface {
	CreateChat(ctx context.Context, chatID string) error
	SaveMessage(ctx context.Context, message models.Message) error
	DeleteMessage(ctx context.Context, chatID string, MessageID uint64) error
	GetMessages(ctx context.Context, chatID string, startIdx, endIdx uint64) ([]models.Message, error)

	Close() error
}

type mongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func (r *mongoRepository) CreateChat(ctx context.Context, chatID string) error {
	slog.Debug("mongo.CreateChat", "chatID", chatID)
	return r.db.CreateCollection(ctx, chatID)
}

func (r *mongoRepository) SaveMessage(ctx context.Context, message models.Message) error {
	slog.Debug("mongo.SaveMessage", "message", message)
	if _, err := r.db.Collection(message.ChatID).InsertOne(ctx, message); err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}
	slog.Debug("saved message", "message", message)
	return nil
}

func (r *mongoRepository) DeleteMessage(ctx context.Context, chatID string, MessageID uint64) error {
	slog.Debug("mongo.DeleteMessage", "chatID", chatID, "MessageID", MessageID)

	// Устанавливаем текущую метку времени для поля DeletedAt
	_, err := r.db.Collection(chatID).UpdateOne(
		ctx,
		bson.M{"_id": MessageID}, // Условие поиска по ID сообщения
		bson.M{
			"$set": bson.M{"deletedat": time.Now()}, // Обновление поля DeletedAt
		},
	)
	if err != nil {
		return fmt.Errorf("failed to logically delete message: %w", err)
	}

	slog.Debug("logically deleted message", "chatID", chatID, "MessageID", MessageID)
	return nil
}

func (r *mongoRepository) GetMessages(
	ctx context.Context,
	chatID string,
	startIdx, endIdx uint64,
) ([]models.Message, error) {
	slog.Debug("mongo.GetMessages", "chatID", chatID, "startIdx", startIdx, "endIdx", endIdx)
	messages := make([]models.Message, 0)
	filter := bson.M{"deletedat": nil}
	findOpts := options.Find()
	findOpts.SetSkip(int64(startIdx))
	findOpts.SetLimit(int64(endIdx - startIdx + 1))

	cursor, err := r.db.Collection(chatID).Find(ctx, filter, findOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var message models.Message
		if err = cursor.Decode(&message); err != nil {
			return nil, fmt.Errorf("failed to decode message: %w", err)
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	slog.Debug("got messages", "chatID", chatID, "messages", messages)
	return messages, nil
}

func (r *mongoRepository) Close() error {
	slog.Info("MongoDB disconnected")
	return r.client.Disconnect(context.Background())
}

func NewMongoRepository(config *config.Mongo) MongoRepository {
	client := mustConnectDB(config)
	slog.Info("MongoDB connected")
	return &mongoRepository{
		client: client,
		db:     client.Database(config.Database),
	}
}

func mustConnectDB(config *config.Mongo) *mongo.Client {
	dsn := getDsn(config)
	opt := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %s", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %s", err)
	}

	slog.Info("MongoDB connected")
	return client
}

func getDsn(config *config.Mongo) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
}
