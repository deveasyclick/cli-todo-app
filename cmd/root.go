/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/yusufniyi/cli-todo-app/cmd/steps"
	"github.com/yusufniyi/cli-todo-app/cmd/ui/textinput"
	"github.com/yusufniyi/cli-todo-app/internal/database"
	"github.com/yusufniyi/cli-todo-app/internal/database/models"
	"github.com/yusufniyi/cli-todo-app/internal/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-todo-app",
	Short: "A lightweight CLI to-do app for managing tasks, setting priorities, and staying organized directly from your terminal.",
	Long: `A simple and efficient CLI-based to-do app that helps you organize tasks, set priorities, and track progress directly from your terminal. 
	
	Perfect for developers and productivity enthusiasts who prefer lightweight tools without distractions.`,
	Run: func(cmd *cobra.Command, args []string) {
		var email, password, name string
		var user models.User
		fmt.Println("Welcome to Easyclick CLI todo app")
		email = readInput("email")
		password = readInput("password")
		userService := &services.UserService{
			DB: database.DB,
		}
		userExists := userService.CheckIfUserAlreadyExistsByEmail(email)
		if !userExists {
			name = readInput("fullname")
			user = userService.Signup(&models.User{Email: email, Password: password, Name: name})
		}

		result := textinput.Result{Output: ""}
		steps := steps.InitSteps()
		step := steps.Steps["actions"]
		p := tea.NewProgram(textinput.InitialTextInputModel(&result, step))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
		fmt.Println("aaaaaaaaaaaa", result.Output, user)
		// create a user if it doesn't exists
		// login a user if it exists
		// Prevent unauthenticated user from running a command
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-todo-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readInput(inputName string) string {
	var input string
	var err error
	subtleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	for {
		fmt.Printf(subtleStyle.Render(fmt.Sprintf("Please enter your %s\n", inputName)))
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}

		if len(strings.TrimSpace(input)) <= 0 {
			continue
		}

		break
	}
	return input
}
