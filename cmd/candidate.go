/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/spf13/cobra"
)

// candidateCmd represents the candidate command
var candidateCmd = &cobra.Command{
	Use:   "candidate",
	Short: "Candidate commands",
	Long:  `Show candidates.`,
	Run: func(cmd *cobra.Command, args []string) {
		list := cmd.Flag("list").Value.String()

		if list == "true" {
			cs := candidate.All()

			for _, c := range cs {
				log.Printf("Name: %s\n", c.Name)
				log.Printf("Party: %s\n", c.Party)
				log.Printf("Code: %s\n", c.Code)

				log.Println()
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
