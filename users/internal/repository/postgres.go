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
	AddUserToGroup(ctx context.Context, ug models.UserGroup) error
	CreateGroup(ctx context.Context, name, groupID string, userIDs []string) error
	DeleteGroup(ctx context.Context, groupID string) error
	GetGroup(ctx context.Context, groupID string) (models.Group, error)
	RemoveUserFromGroup(ctx context.Context, ug models.UserGroup) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (models.User, error)
	UpdateUser(ctx context.Context, user models.UpdateUser) error
	CreateUser(ctx context.Context, user models.User) error
}

type postgres struct {
	db *sql.DB
}

func (p *postgres) RemoveUserFromGroup(ctx context.Context, ug models.UserGroup) error {
	query := `DELETE FROM user_groups WHERE group_id = $1 AND user_id = $2`
	_, err := p.db.ExecContext(ctx, query, ug.GroupID, ug.UserID)
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %w", err)
	}

	return nil
}

func (p *postgres) DeleteUser(ctx context.Context, userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (p *postgres) DeleteGroup(ctx context.Context, groupID string) error {
	query := `DELETE FROM groups WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, groupID)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	return nil
}

func (p *postgres) AddUserToGroup(ctx context.Context, ug models.UserGroup) error {
	query := `INSERT INTO user_groups (group_id, user_id) VALUES ($1, $2)`
	_, err := p.db.ExecContext(ctx, query, ug.GroupID, ug.UserID)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %w", err)
	}

	return nil
}

func (p *postgres) CreateUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (id, name, birth_date) VALUES ($1, $2, $3)`
	_, err := p.db.ExecContext(ctx, query, user.ID, user.Name, user.BirthDate)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (p *postgres) GetUser(ctx context.Context, userID string) (models.User, error) {
	query := `SELECT id, name, birth_date, avatar_url FROM users WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query, userID)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.BirthDate, &user.Avatar); err != nil {
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (p *postgres) UpdateUser(ctx context.Context, user models.UpdateUser) error {
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

func (p *postgres) CreateGroup(ctx context.Context, name, groupID string, userIDs []string) error {
	slog.Debug("postgres.CreateGroup", "name", name, "groupID", groupID, "userIDs", userIDs)

	query := `INSERT INTO groups (id, name) VALUES ($1, $2)`
	_, err := p.db.ExecContext(ctx, query, groupID, name)
	if err != nil {
		return fmt.Errorf("failed to create group: %w", err)
	}

	for _, userID := range userIDs {
		query := `INSERT INTO user_groups (group_id, user_id) VALUES ($1, $2)`
		_, err := p.db.ExecContext(ctx, query, groupID, userID)
		if err != nil {
			return fmt.Errorf("failed to add user to group: %w", err)
		}
	}

	return nil
}

func (p *postgres) GetGroup(ctx context.Context, groupID string) (models.Group, error) {
	query := `SELECT g.id, g.name, g.avatar_url, u.user_id FROM groups g 
              LEFT JOIN user_groups u ON g.id = u.group_id
              WHERE g.id = $1`

	rows, err := p.db.QueryContext(ctx, query, groupID)
	if err != nil {
		return models.Group{}, fmt.Errorf("failed to get group: %w", err)
	}
	defer rows.Close()

	var members []*models.User
	var groupIDOut, groupName string
	var avatar sql.NullString

	for rows.Next() {
		var userID string
		if err := rows.Scan(&groupIDOut, &groupName, &avatar, &userID); err != nil {
			return models.Group{}, fmt.Errorf("failed to get group: %w", err)
		}
		members = append(members, &models.User{ID: userID})
	}

	return models.Group{
		ID:      groupIDOut,
		Name:    groupName,
		Avatar:  &avatar.String,
		Members: members,
	}, nil
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
