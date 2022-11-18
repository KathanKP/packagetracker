/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/KathanKP/packagetracker/internal/domain"
	"github.com/KathanKP/packagetracker/internal/utilities"
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
		newPackage := domain.Package{"FEDEX", "600117021790", "NEW", time.Now()}
		// Call function that adds package to the database table
		err := utilities.AddPackageToDB(newPackage)
		if err != nil {
			log.Fatalln(err)
		}
		// Prints any err, success message
	},
}

func init() {
	rootCmd.AddCommand(addpackageCmd)
}
