package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of task-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task-cli v0.1 -- HEAD")
	},
}
