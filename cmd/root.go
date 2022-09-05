/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "packagetracker",
	Version: version,
	Short:   "An application to track your packages via UPS/Fedex/USPS",
	Long: `This application is used to trakc your packages shipped and 
	delievered via companies like Fedex, UPS and USPS. You can add new packages 
	to track, track the latest status of each package and delete any delivered 
	packaged.`,
	// Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello CLI") },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
