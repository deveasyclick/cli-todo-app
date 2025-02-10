package main

import (
	"github.com/yusufniyi/cli-todo-app/cmd"
	"github.com/yusufniyi/cli-todo-app/internal/db"
)

func main() {
	db.ConnectDatabase()
	defer db.Close()
	cmd.Execute()
}
