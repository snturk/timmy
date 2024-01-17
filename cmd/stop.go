/*
Copyright © 2024 Muratcan Senturk <mcsnturk@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/snturk/timmy/internal/service"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the currently running time entry.",
	Long:  `Stops the currently running time entry. If there is no running time entry, it will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Stopping time tracker... ")
		err := service.StopTimeEntry()
		if err != nil {
			fmt.Errorf("error while stopping time entry: %v", err)
			return
		}
		fmt.Println("Done.")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
