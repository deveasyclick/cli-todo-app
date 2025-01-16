package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	title string
)

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the todo to remove")
	removeCmd.MarkFlagRequired("title")
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Long:  "Remove a todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove command is running", title)
		// Add name to database
	},
}
