package services

import (
	"database/sql"
	"fmt"

	"github.com/yusufniyi/cli-todo-app/internal/database/models"
)

type UserService struct {
	DB *sql.DB
}

func (userService UserService) Login(email string, password string) {
	fmt.Println("Login in")
}

func (userService UserService) Signup(user *models.User) {
	fmt.Println("Sigin up")
}
