package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addTodoTitle string
	description  string
	priority     string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addTodoTitle, "title", "t", "", "Todo title")
	addCmd.Flags().StringVarP(&description, "desc", "d", "", "Todo Description")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "", "Todo priority")
	addCmd.MarkFlagsRequiredTogether("title", "desc")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Long:  "Add a todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command is running", addTodoTitle, description)
		// Add name to database
	},
}
