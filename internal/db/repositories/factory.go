package repositories

import "github.com/yusufniyi/cli-todo-app/internal/db"

type RepositoryFactory interface {
	NewUserRepository() User
	NewTodoRepository() Todo
}

type Factory struct {
}

func (rf *Factory) NewUserRepository() *userRepo {
	return &userRepo{db: db.DBInstance}
}

func (rf *Factory) NewTodoRepository() *todoRepo {
	return &todoRepo{db: db.DBInstance}
}
