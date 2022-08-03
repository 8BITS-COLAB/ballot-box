/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

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
		sk := cmd.Flag("secret-key").Value.String()
		address := cmd.Flag("address").Value.String()
		registry := cmd.Flag("registry").Value.String()
		pvk := cmd.Flag("private-key").Value.String()

		if new != "" && sk != "" {
			voter.New(new, sk)
		}

		if address == "true" {
			v, err := voter.Show(pvk, sk)

			if err != nil {
				log.Fatalf("failed to show voter: %s", err)
			}

			fmt.Printf("Address: %s\n", v.Address)
		}

		if registry == "true" {
			v, err := voter.Show(pvk, sk)

			if err != nil {
				log.Fatalf("failed to show voter: %s", err)
			}

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
	voterCmd.PersistentFlags().StringP("secret-key", "k", "", "voter secret key")
	voterCmd.PersistentFlags().BoolP("address", "a", false, "address voter")
	voterCmd.PersistentFlags().BoolP("registry", "r", false, "registry voter")
	voterCmd.PersistentFlags().StringP("private-key", "p", "", "voter private key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
