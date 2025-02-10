package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
)

type User interface {
	AddUser(user *models.User) (int, error)
	DeleteUser(userId int) error
	FindUser(email string) (models.User, error)
}

type userRepo struct {
	db *pgx.Conn
}

func (ur userRepo) AddUser(user *models.User) (int, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var userID int
	err := ur.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	return userID, nil
}

func (ur userRepo) DeleteUser(userId int) error {
	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	_, err := ur.db.Exec(context.Background(), query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %d: %w", userId, err)
	}

	return nil
}

func (ur userRepo) FindUser(email string) (models.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1;
	`
	user := models.User{}
	err := ur.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_at)

	if err != nil && err != pgx.ErrNoRows {
		return user, fmt.Errorf("Fatal: failed to get user with email %s: %w", email, err)
	}

	return user, nil
}
