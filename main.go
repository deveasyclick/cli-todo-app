/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/yusufniyi/cli-todo-app/cmd"
	"github.com/yusufniyi/cli-todo-app/internal/db"
)

func main() {
	db.ConnectDatabase()
	defer db.DBInstance.Close()
	cmd.Execute()
}
