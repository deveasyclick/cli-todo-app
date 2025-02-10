package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
)

type Todo interface {
	Add(todo *models.Todo) (int, error)
	Delete(todoId int) error
	FindOne(id int) (models.Todo, error)
	Find() ([]models.Todo, error)
}

type todoRepo struct {
	db *pgx.Conn
}

func (tr todoRepo) Add(todo *models.Todo) (int, error) {
	query := `
		INSERT INTO todos (title, description, status, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var todoID int
	err := tr.db.QueryRow(context.Background(), query, todo.Title, todo.Description, todo.Status, todo.UserId).Scan(&todoID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert todo: %w", err)
	}

	return todoID, nil
}

func (tr todoRepo) Delete(todoId int) error {
	query := `
		DELETE FROM todos
		WHERE id = $1;
	`

	_, err := tr.db.Exec(context.Background(), query, todoId)
	if err != nil {
		return fmt.Errorf("failed to delete todo with ID %d: %w", todoId, err)
	}

	return nil
}

func (tr todoRepo) FindOne(id int) (models.Todo, error) {
	query := `
		SELECT id, title, description, status, created_at
		FROM todos
		WHERE id = $1;
	`
	todo := models.Todo{}
	err := tr.db.QueryRow(context.Background(), query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at)

	if err != nil && err != pgx.ErrNoRows {
		return todo, fmt.Errorf("Fatal: failed to get todo with id %d: %w", id, err)
	}

	return todo, nil
}

// Find retrieves all todos from the database and returns them as a slice of models.Todo.
// It returns an error if there is a failure in executing the query or scanning the results.

func (tr todoRepo) Find() ([]models.Todo, error) {
	query := `
		SELECT id, title, description, status, created_at
		FROM todos
	`
	rows, err := tr.db.Query(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("Fatal: failed to fetch todos: %w", err)
	}

	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at)
		if err != nil {
			log.Fatalf("Row scan failed: %v\n", err)
		}
		todos = append(todos, todo)
	}

	if rows.Err() != nil {
		log.Fatalf("Rows error: %v\n", rows.Err())
	}
	return todos, nil
}
