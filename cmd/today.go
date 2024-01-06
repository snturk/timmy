/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/snturk/timmy/internal/service"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your daily time entries",
	Long: `A brief description of your daily time entries. For example:
	Your daily time entries:
	- Task 1: 1h 30m
	- Task 2: 2h 15m
	- Task 3: 3h 45m
	Total: 7h 30m
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.PrintTodayBrief()
		if err != nil {
			fmt.Errorf("error while printing today brief: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
