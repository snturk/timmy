/*
Copyright © 2024 Muratcan Senturk <mcsnturk@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/snturk/timmy/internal/service"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Shows the current time entry, if any.",
	Long:  `Shows the current time entry, if any.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.PrintCurrent()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
