/*
Copyright © 2024 Muratcan Senturk <mcsnturk@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timmy",
	Short: "Simple and easy to use time tracker",
	Long: `Timmy is a simple and easy to use time tracker. You can use it to track, manage and sync your time from any terminal.

More information at https://github.com/snturk/timmy/blob/main/README.md
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
