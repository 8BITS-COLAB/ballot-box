/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/8BITS-COLAB/ballot-box/keystore"
	"github.com/spf13/cobra"
)

// keystoreCmd represents the keystore command
var keystoreCmd = &cobra.Command{
	Use:   "keystore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := keystore.Show()

		fmt.Println(k)
	},
}

func init() {
	rootCmd.AddCommand(keystoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keystoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
