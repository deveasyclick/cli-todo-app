package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/internal/db/models"
	"github.com/yusufniyi/cli-todo-app/internal/db/repositories"
	"github.com/yusufniyi/cli-todo-app/internal/service/auth"
)

var (
	userSignupEmail    string
	userSignupPassword string
	userSignupName     string
)

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.Flags().StringVarP(&userSignupEmail, "email", "e", "", "User email")
	signupCmd.Flags().StringVarP(&userSignupPassword, "password", "p", "", "User Password")
	signupCmd.Flags().StringVarP(&userSignupName, "name", "n", "", "User name")
	signupCmd.MarkFlagsRequiredTogether("email", "password", "name")
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "User Register",
	Long:  "Register your email to have access to the todo commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signup command is running", email, password)
		if userSignupEmail == "" || userSignupPassword == "" || userSignupName == "" {
			log.Fatal("fatal: You must specify the username, email and password")
		}
		auth := auth.AuthService{
			UserRepository: repositories.UserRepository{},
		}
		user := models.User{
			Name:     userSignupName,
			Password: userSignupPassword,
			Email:    userSignupEmail,
		}
		user = auth.Signup(&user)
		fmt.Println(user)
	},
}
