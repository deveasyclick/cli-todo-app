package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	authservice "github.com/yusufniyi/cli-todo-app/internal/service/auth"
	todoservice "github.com/yusufniyi/cli-todo-app/internal/service/todo"
)

var (
	addTodoTitle string
	addTododesc  string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addTodoTitle, "title", "t", "", "Todo title")
	addCmd.Flags().StringVarP(&addTododesc, "desc", "d", "", "Todo Description")
	addCmd.MarkFlagsRequiredTogether("title", "desc")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Long:  "Add a todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command is running", addTodoTitle, addTododesc)
		// Check if user is authenticated
		// Add todo to database
		repoFactory := &repositories.Factory{}
		authService := authservice.New(repoFactory.NewUserRepository())
		token := authService.Authenticate()

		todoService := todoservice.New(repoFactory.NewTodoRepository())
		todoService.AddTodo(addTodoTitle, addTododesc, token.ID)
	},
}
