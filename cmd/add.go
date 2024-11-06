package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}
