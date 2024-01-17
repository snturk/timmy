/*
Copyright © 2024 Muratcan Senturk <mcsnturk@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/snturk/timmy/internal/service"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync local time entry file with Toggl.",
	Long:  `Sync local time entry file with Toggl. Only the entries that are not fetched from Toggl will be fetched.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.FetchTodayToToggl()
		if err != nil {
			fmt.Printf("error while fetching today's time entries to Toggl: %v\n", err)
		} else {
			fmt.Println("successfully fetched today's time entries to Toggl")
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
