package publisher

import (
	"fmt"

	"github.com/alirezazeynali75/nats/internal/config"
	"github.com/alirezazeynali75/nats/internal/logger"
	"github.com/alirezazeynali75/nats/pkg/events"
	"github.com/spf13/cobra"
)

var subject string
var data string

var PubCmd = &cobra.Command{
	Use:   "publisher",
	Short: "A simple NATS publisher",
	Long:  `A simple NATS publisher for sending messages to a subject.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if subject == "" {
			return fmt.Errorf("subject is required")
		}
		if data == "" {
			return fmt.Errorf("data is required")
		}
		return publish(subject, data)
	},
}

func init() {
	// Add flags to accept subject and data as parameters
	PubCmd.Flags().StringVarP(&subject, "subject", "s", "", "Subject to publish to (required)")
	PubCmd.Flags().StringVarP(&data, "data", "d", "", "Data to publish (required)")
	PubCmd.MarkFlagRequired("subject")
	PubCmd.MarkFlagRequired("data")
}

func publish(subject, data string) error {
	// Load configuration
	conf, err := config.Configure()
	if err != nil {
		return err
	}

	// Initialize logger
	logger := logger.GetLogger(conf.App.Env, conf.App.Name)

	// Create and connect NATS client
	natsClient := events.New(conf.Nats.Url, logger)
	err = natsClient.Connect()
	if err != nil {
		return err
	}
	defer natsClient.Disconnect()

	// Publish the message
	err = natsClient.Publish(subject, []byte(data))
	if err != nil {
		return err
	}

	logger.Info("Message published successfully", "subject", subject, "data", data)
	return nil
}
