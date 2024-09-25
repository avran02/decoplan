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
	AddUserToGroup(ctx context.Context, userGroup models.UserGroup) error
	CreateGroup(ctx context.Context, name string, userIDs []string) (string, error)
	DeleteGroup(ctx context.Context, groupID string) error
	GetGroup(ctx context.Context, groupID string) (models.Group, error)
	RemoveUserFromGroup(ctx context.Context, userGroup models.UserGroup) error
	CreateUser(ctx context.Context, id, name string, birthDate time.Time) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUser) error
}

type userService struct {
	repo repository.Repository
}

func (s *userService) AddUserToGroup(ctx context.Context, userGroup models.UserGroup) error {
	return s.repo.AddUserToGroup(ctx, userGroup)
}

func (s *userService) CreateGroup(ctx context.Context, name string, userIDs []string) (string, error) {
	groupID := uuid.NewString()
	if err := s.repo.CreateGroup(ctx, name, groupID, userIDs); err != nil {
		return "", fmt.Errorf("failed to create group: %w", err)
	}

	return groupID, nil
}

func (s *userService) DeleteGroup(ctx context.Context, groupID string) error {
	return s.repo.DeleteGroup(ctx, groupID)

}

func (s *userService) GetGroup(ctx context.Context, groupID string) (models.Group, error) {
	return s.repo.GetGroup(ctx, groupID)
}

func (s *userService) RemoveUserFromGroup(ctx context.Context, userGroup models.UserGroup) error {
	return s.repo.RemoveUserFromGroup(ctx, userGroup)

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
