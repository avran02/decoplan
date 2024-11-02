package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/avran02/decoplan/chat-storage/internal/config"
	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	CacheGroupLastMessages(ctx context.Context, chatID string, messages []models.Message) error
	GetCacheLimits(ctx context.Context, chatID string) (uint64, uint64, error)

	SaveMessage(ctx context.Context, message models.Message) error
	GetMessages(ctx context.Context, chatID string, startIdx, endIdx uint64) ([]models.Message, error)
	DeleteMessage(ctx context.Context, chatID string, messageID uint64) error

	Close() error
}

type redisRepository struct {
	db *redis.Client
}

func (r *redisRepository) CacheGroupLastMessages(
	ctx context.Context,
	chatID string,
	messages []models.Message,
) error {
	slog.Debug("redis.CacheGroupLastMessages", "chatID", chatID, "messages", messages)

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].ID > messages[j].ID
	})

	startCacheIdx := messages[0].ID
	endCacheIdx := messages[len(messages)-1].ID
	r.setCacheLimits(ctx, chatID, startCacheIdx, endCacheIdx)

	for _, message := range messages {
		if err := r.SaveMessage(ctx, message); err != nil {
			return fmt.Errorf("failed to save message: %w", err)
		}
	}

	slog.Debug("cached messages", "chatID", chatID, "messages", messages)
	return nil
}

func (r *redisRepository) SaveMessage(ctx context.Context, message models.Message) error {
	slog.Debug("redis.SaveMessage", "message", message)

	key := fmt.Sprintf("%s:%d", message.ChatID, message.ID)
	if err := r.db.Set(ctx, key, message, 0).Err(); err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}

	start, end, err := r.GetCacheLimits(ctx, message.ChatID)
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("failed to get cache limits: %w", err)
	}
	if errors.Is(err, redis.Nil) {
		r.setCacheLimits(ctx, message.ChatID, message.ID, message.ID)
		slog.Debug("set cache limits", "chatID", message.ChatID, "startCacheIdx", message.ID, "endCacheIdx", message.ID)
		return nil
	}

	r.setCacheLimits(ctx, message.ChatID, start, end+1)

	slog.Debug("saved message", "message", message)
	return nil
}

func (r *redisRepository) GetCacheLimits(ctx context.Context, chatID string) (uint64, uint64, error) {
	slog.Debug("redis.GetCacheLimits", "chatID", chatID)

	cacheLimits, err := r.db.Get(ctx, fmt.Sprintf("limit:%s", chatID)).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get cache limits: %w", err)
	}

	parts := strings.Split(cacheLimits, ":")
	startIdx, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get cache limits: %w", err)
	}
	endIdx, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get cache limits: %w", err)
	}

	slog.Debug("got cache limits", "chatID", chatID, "startIdx", startIdx, "endIdx", endIdx)
	return startIdx, endIdx, nil
}

func (r *redisRepository) GetMessages(
	ctx context.Context,
	chatID string,
	startIdx, endIdx uint64,
) ([]models.Message, error) {
	slog.Debug("redis.GetMessages", "chatID", chatID, "startIdx", startIdx, "endIdx", endIdx)
	messages := make([]models.Message, 0, endIdx-startIdx+1)
	for id := startIdx; id <= endIdx; id++ {
		key := fmt.Sprintf("%s:%d", chatID, id)
		slog.Debug("getting message", "key", key)
		message := models.Message{}
		err := r.db.Get(ctx, key).Scan(&message)
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("failed to get messages: %w", err)
		}
		if message.DeletedAt != nil || errors.Is(err, redis.Nil) {
			continue
		}
		messages = append(messages, message)
	}

	slog.Debug("got messages", "chatID", chatID, "messages", messages)
	return messages, nil
}

func (r *redisRepository) DeleteMessage(ctx context.Context, chatID string, messageID uint64) error {
	slog.Debug("redis.DeleteMessage", "chatID", chatID, "messageID", messageID)
	key := fmt.Sprintf("%s:%d", chatID, messageID)
	message := models.Message{}
	err := r.db.Get(ctx, key).Scan(&message)
	if err != nil {
		return fmt.Errorf("failed to get message: %w", err)
	}
	currentTime := time.Now()
	message.DeletedAt = &currentTime

	if err = r.db.Set(ctx, key, message, 0).Err(); err != nil {
		return fmt.Errorf("failed to logically delete message: %w", err)
	}

	slog.Debug("deleted message", "chatID", chatID, "messageID", messageID)
	return nil
}

func (r *redisRepository) Close() error {
	return r.db.Close()
}

func (r *redisRepository) setCacheLimits(ctx context.Context, chatID string, startCacheIdx, endCacheIdx uint64) error {
	slog.Debug("redis.setCacheLimits", "chatID", chatID, "startCacheIdx", startCacheIdx, "endCacheIdx", endCacheIdx)
	key := fmt.Sprintf("limit:%s", chatID)
	val := fmt.Sprintf("%d:%d", startCacheIdx, endCacheIdx)

	if err := r.db.Set(ctx, key, val, 0).Err(); err != nil {
		return fmt.Errorf("failed to set cache limits: %w", err)
	}

	slog.Debug("set cache limits", "chatID", chatID, "startCacheIdx", startCacheIdx, "endCacheIdx", endCacheIdx)
	return nil
}

func parseLimits(limits string) (uint64, uint64, error) {
	slog.Debug("parseLimits", "limits", limits)
	limitsList := strings.Split(limits, ":")
	if len(limitsList) != 2 {
		return 0, 0, ErrBrokenLimits
	}
	startCacheIdx, err := strconv.ParseUint(limitsList[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Broken limits: %w", err)
	}
	endCacheIdx, err := strconv.ParseUint(limitsList[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Broken limits: %w", err)
	}
	slog.Debug("parseLimits parsed", "startCacheIdx", startCacheIdx, "endCacheIdx", endCacheIdx)
	return startCacheIdx, endCacheIdx, nil
}

func NewRedisRepository(config *config.Redis) RedisRepository {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	slog.Info("Connected to redis")
	return &redisRepository{
		db: db,
	}
}
