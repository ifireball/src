/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load repo configuration and apply to the src directory",
	Long: `Load the repository configuration including remotes and branches\n
and generate an src directory with clones of all those repos and with\n
a similar branch configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("load called")
	},
}

func init() {
	configCmd.AddCommand(loadCmd)
	loadCmd.Flags().Bool("dry-run", false, "Print what would be done but don't apply changes")
}
