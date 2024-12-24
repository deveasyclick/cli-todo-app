package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var yusufCmd = &cobra.Command{
	Use:   "yusuf",
	Short: "Print log by yusuf",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yusuf testing cli")
	},
}
