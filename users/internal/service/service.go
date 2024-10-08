package service

import (
	"context"
	"fmt"
	"time"

	"github.com/avran02/decoplan/users/internal/models"
	"github.com/avran02/decoplan/users/internal/repository"
	"github.com/google/uuid"
)

type UserService interface {
	AddUserToChat(ctx context.Context, userChat models.UserChat) error
	CreateChat(ctx context.Context, name string, userIDs []string) (string, error)
	DeleteChat(ctx context.Context, chatID string) error
	GetChat(ctx context.Context, chatID string) (models.Chat, error)
	RemoveUserFromChat(ctx context.Context, userChat models.UserChat) error
	CreateUser(ctx context.Context, id, name string, birthDate time.Time) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUser) error
}

type userService struct {
	repo repository.Repository
}

func (s *userService) AddUserToChat(ctx context.Context, userChat models.UserChat) error {
	return s.repo.AddUserToChat(ctx, userChat)
}

func (s *userService) CreateChat(ctx context.Context, name string, userIDs []string) (string, error) {
	chatID := uuid.NewString()
	if err := s.repo.CreateChat(ctx, name, chatID, userIDs); err != nil {
		return "", fmt.Errorf("failed to create chat: %w", err)
	}

	return chatID, nil
}

func (s *userService) DeleteChat(ctx context.Context, chatID string) error {
	return s.repo.DeleteChat(ctx, chatID)

}

func (s *userService) GetChat(ctx context.Context, chatID string) (models.Chat, error) {
	return s.repo.GetChat(ctx, chatID)
}

func (s *userService) RemoveUserFromChat(ctx context.Context, userChat models.UserChat) error {
	return s.repo.RemoveUserFromChat(ctx, userChat)

}
func (s *userService) CreateUser(ctx context.Context, id, name string, birthDate time.Time) error {
	user := models.User{
		ID:        id,
		Name:      name,
		BirthDate: birthDate,
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, userID string) (models.User, error) {
	user, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user models.UpdateUser) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	return s.repo.DeleteUser(ctx, userID)
}

func New(repo repository.Repository) UserService {
	return &userService{repo: repo}
}
