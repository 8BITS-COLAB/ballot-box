/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/spf13/cobra"
)

// candidateCmd represents the candidate command
var candidateCmd = &cobra.Command{
	Use:   "candidate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list := cmd.Flag("list").Value.String()

		if list == "true" {
			cs := candidate.All()

			for _, c := range cs {
				fmt.Printf("Name: %s\n", c.Name)
				fmt.Printf("Party: %s\n", c.Party)
				fmt.Printf("Code: %s\n", c.Code)

				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(candidateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	candidateCmd.PersistentFlags().BoolP("list", "l", false, "list candidates")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// candidateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
