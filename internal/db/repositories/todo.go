package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/yusufniyi/cli-todo-app/internal/db"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
)

type Todo struct {
}

func (todoRepository Todo) Add(todo *models.Todo) (int, error) {
	query := `
		INSERT INTO todos (title, description, status, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var todoID int
	err := db.DBInstance.Conn.QueryRow(context.Background(), query, todo.Title, todo.Description, todo.Status, todo.UserId).Scan(&todoID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert todo: %w", err)
	}

	return todoID, nil
}

func (todoRepository Todo) Delete(todoId int) error {
	query := `
		DELETE FROM todos
		WHERE id = $1;
	`

	_, err := db.DBInstance.Conn.Exec(context.Background(), query, todoId)
	if err != nil {
		return fmt.Errorf("failed to delete todo with ID %d: %w", todoId, err)
	}

	return nil
}

func (todoRepository Todo) FindOne(id int) (models.Todo, error) {
	query := `
		SELECT id, title, description, status, created_at
		FROM todos
		WHERE id = $1;
	`
	todo := models.Todo{}
	err := db.DBInstance.Conn.QueryRow(context.Background(), query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Created_at)

	if err != nil && err != pgx.ErrNoRows {
		return todo, fmt.Errorf("Fatal: failed to get todo with id %d: %w", id, err)
	}

	return todo, nil
}

func (todoRepository Todo) Find() ([]models.Todo, error) {
	query := `
		SELECT id, title, description, status, created_at
		FROM todos
	`
	rows, err := db.DBInstance.Conn.Query(context.Background(), query)

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
