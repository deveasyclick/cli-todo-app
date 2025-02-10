package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	authservice "github.com/yusufniyi/cli-todo-app/internal/service/auth"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  "Login of todo app",
	Run: func(cmd *cobra.Command, args []string) {
		repoFactory := &repositories.Factory{}
		authService := authservice.New(repoFactory.NewUserRepository())
		authService.Logout()
	},
}
