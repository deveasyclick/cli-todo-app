package todoService

import (
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
)

type TodoService struct {
	Repository repositories.Todo
}

func (todoService *TodoService) AddTodo(title string, desc string, userId int) {
	todoObject := models.Todo{
		Title:       title,
		Description: desc,
		UserId:      userId,
		Status:      string(models.InProgress),
	}
	_, err := todoService.Repository.Add(&todoObject)
	var pgErr *pgconn.PgError
	if err != nil && errors.As(err, &pgErr) && pgErr.Code == "23505" {
		log.Fatalf("Fatal: Todo with title %s already exists", title)
	} else if err != nil {
		log.Fatalf("Fatal: Unable to add todo with title %s to database: %s", title, err)
	}

	log.Printf("Todo %s added to the database\n", title)
}

func (todoService *TodoService) RemoveTodo(title string) {

}

func (todoService *TodoService) ListTodos() {

}
