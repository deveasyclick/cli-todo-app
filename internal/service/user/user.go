package user_service

import "github.com/jackc/pgx/v5"

type UserService struct {
	DB pgx.Conn
}

func (userService *UserService) Login() {

}
