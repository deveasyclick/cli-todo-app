package todoservice

import (
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
)

type service struct {
	todoRepository repositories.Todo
}

func (todoService *service) AddTodo(title string, desc string, userId int) {
	todoObject := models.Todo{
		Title:       title,
		Description: desc,
		UserId:      userId,
		Status:      string(models.InProgress),
	}
	_, err := todoService.todoRepository.Add(&todoObject)
	var pgErr *pgconn.PgError
	// Log different message for duplicate todo error
	if err != nil && errors.As(err, &pgErr) && pgErr.Code == "23505" {
		log.Fatalf("Fatal: Todo with title %s already exists", title)
	} else if err != nil {
		log.Fatalf("Fatal: Unable to add todo with title %s to database: %s", title, err)
	}

	log.Printf("Todo %s added to the database\n", title)
}

func (todoService *service) RemoveTodo(title string) {

}

func (todoService *service) ListTodos() {
	todos, err := todoService.todoRepository.Find()
	if err != nil {
		log.Fatalf("Fatal: Unable to fetch todos from database: %s", err)
	}

	// Create a table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Description", "Status"})

	for _, value := range todos {
		t.AppendRow(table.Row{value.ID, value.Title, value.Description, value.Status})
	}
	t.Render()
}

func New(todoRepository repositories.Todo) *service {
	return &service{
		todoRepository: todoRepository,
	}
}
