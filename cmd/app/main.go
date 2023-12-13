package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.Use = "app"
	rootCmd.Short = "It starts as a GraphQL server"
	rootCmd.Version = "1.0"
	rootCmd.SilenceUsage = true

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runApp()
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
