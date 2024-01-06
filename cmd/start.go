/*
Copyright © 2024 Muratcan Senturk <mcsnturk@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/snturk/timmy/internal/service"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the time tracker",
	Long:  `This will start the time tracker with provided task name.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Task name is required
		if len(args) < 1 {
			fmt.Println("Task name is required")
			return
		}

		// Get task name from user
		taskName := args[0]

		err := service.StartTimeEntry(taskName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Starting time tracker for task: ", taskName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Get task name from user
	startCmd.Flags().StringP("task", "t", "", "Task name")
}
