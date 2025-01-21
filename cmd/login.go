package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	"github.com/yusufniyi/cli-todo-app/internal/service/auth"
)

var (
	email    string
	password string
)

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "User Password")
	loginCmd.MarkFlagsRequiredTogether("email", "password")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to app",
	Long:  "Login to app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login command is running", email, password)
		if email == "" || password == "" {
			log.Fatal("fatal: You must specify the usernname and password")
		}
		auth := auth.AuthService{
			UserRepository: repositories.UserRepository{},
		}
		auth.Login(email, password)
	},
}
