/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  `Serve Command`,
	Run: func(cmd *cobra.Command, args []string) {
		inputValue, _ := cmd.Flags().GetString("sampleFlag")

		fmt.Println(inputValue)
		fmt.Println("command called")
		_, err := ping.NewPinger(inputValue)
		if err != nil {
			fmt.Println("Error")
		} else {
			fmt.Println("No Error")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().String("sampleFlag", "www.google.com", "Description for Flag")
}
