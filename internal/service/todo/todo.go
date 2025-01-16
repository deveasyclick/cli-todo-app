package todo_service

import "github.com/jackc/pgx/v5"

type TodoService struct {
	DB pgx.Conn
}

func (todoService *TodoService) AddTodo(title string, desc string) {

}

func (todoService *TodoService) RemoveTodo(title string) {

}

func (todoService *TodoService) ListTodos() {

}
