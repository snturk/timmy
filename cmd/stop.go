/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
