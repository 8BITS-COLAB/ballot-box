/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/8BITS-COLAB/ballot-box/voter"
	"github.com/spf13/cobra"
)

// voterCmd represents the voter command
var voterCmd = &cobra.Command{
	Use:   "voter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		new := cmd.Flag("new").Value.String()
		address := cmd.Flag("address").Value.String()
		registry := cmd.Flag("registry").Value.String()

		if new != "" {
			voter.New(new)
		}

		if address == "true" {
			v := voter.Show()

			fmt.Printf("Address: %s\n", v.Address)
		}

		if registry == "true" {
			v := voter.Show()

			fmt.Printf("Registry: %s\n", v.Registry)
		}
	},
}

func init() {
	rootCmd.AddCommand(voterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	voterCmd.PersistentFlags().StringP("new", "n", "", "create new voter by registry")
	voterCmd.PersistentFlags().BoolP("address", "a", false, "address voter")
	voterCmd.PersistentFlags().BoolP("registry", "r", false, "registry voter")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
