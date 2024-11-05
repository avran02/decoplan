package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"strings"

	_ "github.com/lib/pq"

	"github.com/avran02/decoplan/users/internal/config"
	"github.com/avran02/decoplan/users/internal/models"
)

type Repository interface {
	AddUserToChat(ctx context.Context, ug models.UserChat) error
	CreateChat(ctx context.Context, name, chatID string, userIDs []string) error
	DeleteChat(ctx context.Context, chatID string) error
	GetChat(ctx context.Context, chatID string) (*models.Chat, error)
	RemoveUserFromChat(ctx context.Context, ug models.UserChat) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUser) error
	CreateUser(ctx context.Context, user models.User) error
	GetUserChats(ctx context.Context, userID string) ([]models.Chat, error)
}

type postgres struct {
	db *sql.DB
}

func (p *postgres) RemoveUserFromChat(ctx context.Context, ug models.UserChat) error {
	slog.Info("repository.RemoveUserFromChat")
	slog.Debug("args", "ug", ug)
	query := `DELETE FROM user_chats WHERE chat_id = $1 AND user_id = $2`
	_, err := p.db.ExecContext(ctx, query, ug.ChatID, ug.UserID)
	if err != nil {
		return fmt.Errorf("failed to remove user from chat: %w", err)
	}

	return nil
}

func (p *postgres) DeleteUser(ctx context.Context, userID string) error {
	slog.Info("repository.DeleteUser")
	slog.Debug("args", "userID", userID)
	query := `DELETE FROM users WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (p *postgres) DeleteChat(ctx context.Context, chatID string) error {
	slog.Info("repository.DeleteChat")
	slog.Debug("args", "chatID", chatID)
	query := `DELETE FROM chats WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, chatID)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}

func (p *postgres) AddUserToChat(ctx context.Context, ug models.UserChat) error {
	slog.Info("repository.AddUserToChat")
	slog.Debug("args", "ug", ug)
	query := `INSERT INTO user_chats (chat_id, user_id) VALUES ($1, $2)`
	_, err := p.db.ExecContext(ctx, query, ug.ChatID, ug.UserID)
	if err != nil {
		return fmt.Errorf("failed to add user to chat: %w", err)
	}

	return nil
}

func (p *postgres) CreateUser(ctx context.Context, user models.User) error {
	slog.Info("repository.CreateUser")
	slog.Debug("args", "user", user)
	query := `INSERT INTO users (id, name, birth_date) VALUES ($1, $2, $3)`
	_, err := p.db.ExecContext(ctx, query, user.ID, user.Name, user.BirthDate)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (p *postgres) GetUser(ctx context.Context, userID string) (models.User, error) {
	slog.Info("repository.GetUser")
	slog.Debug("args", "userID", userID)
	query := `SELECT id, name, birth_date, avatar_url FROM users WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query, userID)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.BirthDate, &user.Avatar); err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (p *postgres) UpdateUser(ctx context.Context, user models.UpdateUser) error {
	slog.Info("repository.UpdateUser")
	slog.Debug("args", "user", user)
	var setParts []string
	var args []interface{}
	var argPos int = 1

	if user.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argPos))
		args = append(args, user.Name)
		argPos++
	}

	if !user.BirthDate.IsZero() {
		setParts = append(setParts, fmt.Sprintf("birth_date = $%d", argPos))
		args = append(args, user.BirthDate)
		argPos++
	}

	if user.Avatar != nil {
		setParts = append(setParts, fmt.Sprintf("avatar_url = $%d", argPos))
		args = append(args, *user.Avatar)
		argPos++
	}

	if len(setParts) == 0 {
		return ErrNothingToUpdate
	}

	args = append(args, user.ID)

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $%d`, strings.Join(setParts, ", "), argPos)

	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (p *postgres) CreateChat(ctx context.Context, name, chatID string, userIDs []string) error {
	slog.Debug("postgres.CreateChat", "name", name, "chatID", chatID, "userIDs", userIDs)

	query := `INSERT INTO chats (id, name) VALUES ($1, $2)`
	_, err := p.db.ExecContext(ctx, query, chatID, name)
	if err != nil {
		return fmt.Errorf("failed to create chat: %w", err)
	}

	for _, userID := range userIDs {
		query := `INSERT INTO user_chats (chat_id, user_id) VALUES ($1, $2)`
		_, err := p.db.ExecContext(ctx, query, chatID, userID)
		if err != nil {
			return fmt.Errorf("failed to add user to chat: %w", err)
		}
	}

	return nil
}

func (p *postgres) GetChat(ctx context.Context, chatID string) (*models.Chat, error) {
	slog.Info("repository.GetChat")
	slog.Debug("args", "chatID", chatID)
	query := `SELECT c.id, c.name, c.avatar_url, u.user_id FROM chats c 
              JOIN user_chats u ON c.id = u.chat_id
              WHERE c.id = $1`

	rows, err := p.db.QueryContext(ctx, query, chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}
	defer rows.Close()

	var members []*models.User
	var chatIDOut, chatName string
	var avatar sql.NullString
	numRows := 0
	for rows.Next() {
		numRows++
		var userID string
		if err := rows.Scan(&chatIDOut, &chatName, &avatar, &userID); err != nil {
			return nil, fmt.Errorf("failed to get chat: %w", err)
		}
		members = append(members, &models.User{ID: userID})
	}
	if numRows == 0 {
		return nil, sql.ErrNoRows
	}

	return &models.Chat{
		ID:      chatIDOut,
		Name:    chatName,
		Avatar:  &avatar.String,
		Members: members,
	}, nil
}

func (p *postgres) GetUserChats(ctx context.Context, userID string) ([]models.Chat, error) {
	slog.Info("repository.GetUserChats")
	slog.Debug("args", "userID", userID)
	query := `SELECT c.id, c.name, c.avatar_url FROM chats c 
			  JOIN user_chats u ON c.id = u.chat_id
			  WHERE u.user_id = $1`

	rows, err := p.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user chats: %w", err)
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chatID, chatName string
		var avatar sql.NullString
		if err := rows.Scan(&chatID, &chatName, &avatar); err != nil {
			return nil, fmt.Errorf("failed to get user chats: %w", err)
		}
		chats = append(chats, models.Chat{
			ID:      chatID,
			Name:    chatName,
			Avatar:  &avatar.String,
			Members: nil,
		})
	}

	return chats, nil
}

func New(conf config.DB) Repository {
	db, err := sql.Open("postgres", getDsn(conf))
	if err != nil {
		log.Fatal("can't open db conn:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("can't ping:", err)
	}
	return &postgres{
		db: db,
	}
}

func getDsn(conf config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
	)
}
