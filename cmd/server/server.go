/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"github.com/spf13/cobra"
	"github.com/the-mold/go-project-layout/cmd"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(serverCmd)
}
