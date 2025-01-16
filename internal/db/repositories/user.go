package repositories

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
)

type UserRepository struct {
	DB *pgx.Conn
}

func (userRepository UserRepository) AddUser(user *models.User) (int, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var userID int
	err := userRepository.DB.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	return userID, nil
}

func (userRepository UserRepository) DeleteUser(userId int) error {
	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	_, err := userRepository.DB.Exec(context.Background(), query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %d: %w", userId, err)
	}

	return nil
}

func (userRepository UserRepository) FindUser(email string) (models.User, error) {
	query := `
		SELECT id, name, email, createdAt
		FROM users
		WHERE email = $1;
	`
	var user models.User
	err := userRepository.DB.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Created_at)

	if err != nil {
		return user, fmt.Errorf("failed to get user with email %s: %w", email, err)
	}

	return user, nil
}
