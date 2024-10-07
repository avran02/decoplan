package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/avran02/decoplan/chat-storage/internal/models"
	"github.com/avran02/decoplan/chat-storage/internal/repository"
)

type Service interface {
	SaveMessage(ctx context.Context, message models.Message) error
	GetMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error)
	DeleteMessage(ctx context.Context, chatID string, messageID uint64) error
	CacheLastMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error)
	CreateChat(ctx context.Context, chatID string) error
}

type service struct {
	redis repository.RedisRepository
	mongo repository.MongoRepository
}

func (s *service) SaveMessage(ctx context.Context, message models.Message) error {
	slog.Debug("service.SaveMessage", "message", message)
	s.mongo.SaveMessage(ctx, message)
	return s.redis.SaveMessage(ctx, message)
}

func (s *service) GetMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error) {
	// todo: add more logs
	slog.Debug("service.GetMessages", "chatID", chatID, "limit", limit, "offset", offset)
	var messages []models.Message
	cacheStartIdx, cacheEndIdx, err := s.redis.GetCacheLimits(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cache limits: %w", err)
	}
	requestingStartIdx := offset
	requestingEndIdx := limit + offset - 1

	slog.Debug(
		"service.SaveMessage: choosing between redis and mongo",
		"requestingStartIdx", requestingStartIdx,
		"requestingEndIdx", requestingEndIdx,
		"cacheStartIdx", cacheStartIdx,
		"cacheEndIdx", cacheEndIdx,
	)
	if requestingStartIdx >= cacheStartIdx && requestingEndIdx <= cacheEndIdx {
		return s.redis.GetMessages(ctx, chatID, requestingStartIdx, requestingEndIdx)
	}

	// вообще, тут должны быть строгие знаки, чтобы границы тоже брать из кэша
	// но это слишком много дополнительного усложнения, ради двух сообщений
	if requestingStartIdx >= cacheEndIdx || requestingEndIdx <= cacheStartIdx {
		return s.mongo.GetMessages(ctx, chatID, requestingStartIdx, requestingEndIdx)
	}

	if requestingStartIdx < cacheStartIdx {
		cachedMessages, err := s.redis.GetMessages(ctx, chatID, cacheStartIdx, requestingEndIdx)
		if err != nil {
			return nil, fmt.Errorf("failed to get messages: %w", err)
		}
		mongoMessages, err := s.mongo.GetMessages(ctx, chatID, requestingStartIdx+1, cacheStartIdx)
		if err != nil {
			return nil, fmt.Errorf("failed to get messages: %w", err)
		}
		messages = append(messages, mongoMessages...)
		messages = append(messages, cachedMessages...)
		return messages, nil
	}

	if requestingEndIdx > cacheEndIdx {
		cachedMessages, err := s.redis.GetMessages(ctx, chatID, requestingStartIdx, cacheEndIdx)
		if err != nil {
			return nil, fmt.Errorf("failed to get messages: %w", err)
		}
		mongoMessages, err := s.mongo.GetMessages(ctx, chatID, cacheEndIdx+1, requestingEndIdx)
		if err != nil {
			return nil, fmt.Errorf("failed to get messages: %w", err)
		}
		messages = append(messages, cachedMessages...)
		messages = append(messages, mongoMessages...)
		return messages, nil
	}

	return nil, ErrUnexpectedBehavior
}

func (s *service) DeleteMessage(ctx context.Context, chatID string, messageID uint64) error {
	slog.Debug("service.DeleteMessage", "chatID", chatID, "messageID", messageID)
	s.mongo.DeleteMessage(ctx, chatID, messageID)
	return s.redis.DeleteMessage(ctx, chatID, messageID)
}

func (s *service) CacheLastMessages(ctx context.Context, chatID string, limit, offset uint64) ([]models.Message, error) {
	slog.Debug("service.CacheLastMessages", "chatID", chatID, "limit", limit, "offset", offset)
	messages, err := s.mongo.GetMessages(ctx, chatID, offset, offset+limit-1)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	return messages, s.redis.CacheGroupLastMessages(ctx, chatID, messages)
}

func (s *service) CreateChat(ctx context.Context, chatID string) error {
	slog.Debug("service.CreateChat", "chatID", chatID)
	return s.mongo.CreateChat(ctx, chatID)
}

func New(m repository.MongoRepository, r repository.RedisRepository) Service {
	slog.Info("initializing service")
	return &service{
		mongo: m,
		redis: r,
	}
}
