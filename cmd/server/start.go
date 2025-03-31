/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		isDetailed, err := cmd.Flags().GetBool("detailed")
		if err != nil {
			return err
		}

		if isDetailed {
			fmt.Println("server start called with detailed output")
		} else {
			fmt.Println("server start called")
		}
		return nil
	},
}

func init() {
	startCmd.Flags().BoolP("detailed", "d", false, "Help message for the flag")

	serverCmd.AddCommand(startCmd)
}
