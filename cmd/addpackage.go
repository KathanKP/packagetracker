/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addpackageCmd represents the addpackage command
var addpackageCmd = &cobra.Command{
	Use:   "addpackage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addpackage called")
		// Take as input package details
		// This calls the server API func to add package based on the the delivery service
		// Prints any err, success message
	},
}

func init() {
	rootCmd.AddCommand(addpackageCmd)
}
