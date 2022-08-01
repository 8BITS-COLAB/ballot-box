/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/8BITS-COLAB/ballot-box/vote"
	"github.com/spf13/cobra"
)

// voteCmd represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		code := cmd.Flag("code").Value.String()
		key := cmd.Flag("key").Value.String()
		status := cmd.Flag("status").Value.String()

		if code != "" {
			v := vote.New(code, key)

			fmt.Printf("vote %d added\n", v.Index)
		}

		if status == "true" {
			vts := vote.Status()

			for key, value := range vts {
				fmt.Printf("Candidate %s have %d votes\n", key, value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	voteCmd.PersistentFlags().StringP("code", "c", "", "candidate code")
	voteCmd.PersistentFlags().BoolP("status", "s", false, "votes status")
	voteCmd.PersistentFlags().StringP("key", "k", "", "voter key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
