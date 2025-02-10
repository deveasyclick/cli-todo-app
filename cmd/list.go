package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	authservice "github.com/yusufniyi/cli-todo-app/internal/service/auth"
	todoservice "github.com/yusufniyi/cli-todo-app/internal/service/todo"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  "List todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command is running", title)
		// Add name to database
		repoFactory := &repositories.Factory{}
		authService := authservice.New(repoFactory.NewUserRepository())
		authService.Authenticate()

		todoService := todoservice.New(repoFactory.NewTodoRepository())
		todoService.ListTodos()
	},
}
