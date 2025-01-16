package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	},
}
