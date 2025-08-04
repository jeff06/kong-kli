package cmd

import (
	"github.com/spf13/cobra"
	"kong-kli/cmd/consumer_groups"
	"kong-kli/cmd/consumers"
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

	rootCmd.PersistentFlags().StringP("controlPlaneId", "c", "", "The id of the control plane targeted")
	rootCmd.MarkFlagRequired("controlPlaneId")

	rootCmd.AddCommand(consumers.CmdConsumer)
	rootCmd.AddCommand(consumer_groups.CmdConsumerGroups)
}
