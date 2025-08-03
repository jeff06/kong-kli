package cmd

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumer"
)

var (
	rootCmd = &cobra.Command{
		Use:   "kong-kli",
		Short: "A CLI for managing Kong API Gateway",
		Long: `Kong-kli is a command-line interface for managing Kong API Gateway.
It provides commands for working with various Kong resources
such as services, routes, consumers, and plugins.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Kong token to use for authentication")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.PersistentFlags().StringP("server", "s", "", "Kong Konnect Server")
	rootCmd.MarkPersistentFlagRequired("server")

	rootCmd.AddCommand(consumer.CmdConsumer)
}
