/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/8BITS-COLAB/ballot-box/env"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		set := cmd.Flag("set").Value.String()
		get := cmd.Flag("get").Value.String()

		if set != "" {
			v := strings.Split(set, ":=")

			env.Set(v[0], v[1])

		}

		if get != "" {
			v, err := env.Get(get)

			if err != nil {
				log.Fatalf("failed to get config: %s", err)
			}

			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	configCmd.PersistentFlags().StringP("set", "s", "", "set config")
	configCmd.PersistentFlags().StringP("get", "g", "", "get config")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
