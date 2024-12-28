/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/yusufniyi/cli-todo-app/cmd"
	"github.com/yusufniyi/cli-todo-app/internal/database"
	"github.com/yusufniyi/cli-todo-app/internal/loaders"
)

func main() {
	loaders.LoadEnvVariables()
	database.Open()
	defer database.Close()
	cmd.Execute()
}
