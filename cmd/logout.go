package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	"github.com/yusufniyi/cli-todo-app/internal/service/auth"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  "Login of todo app",
	Run: func(cmd *cobra.Command, args []string) {
		auth := auth.AuthService{
			UserRepository: repositories.UserRepository{},
		}
		auth.Logout()
	},
}
