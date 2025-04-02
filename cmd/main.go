package main

import (
	"github.com/alirezazeynali75/nats/cmd/publisher"
	"github.com/alirezazeynali75/nats/cmd/subscriber"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nats-client",
	Short: "A simple NATS client",
	Long:  `A simple NATS client for publishing and subscribing to messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(publisher.PubCmd)
	rootCmd.AddCommand(subscriber.SubCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

