package subscriber

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alirezazeynali75/nats/internal/config"
	"github.com/alirezazeynali75/nats/internal/logger"
	"github.com/alirezazeynali75/nats/pkg/events"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var subject string

var SubCmd = &cobra.Command{
	Use:   "subscriber",
	Short: "A simple NATS subscriber",
	Long:  `A simple NATS subscriber for receiving messages from a subject.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if subject == "" {
			return fmt.Errorf("subject is required")
		}
		return start(context.Background(), subject)
	},
}

func init() {
	// Add a flag to accept the subject as a parameter
	SubCmd.Flags().StringVarP(&subject, "subject", "s", "", "Subject to subscribe to (required)")
	SubCmd.MarkFlagRequired("subject")
}

func start(ctx context.Context, subject string) error {
	conf, err := config.Configure()
	if err != nil {
		return err
	}

	logger := logger.GetLogger(conf.App.Env, conf.App.Name)
	natsClient := events.New(conf.Nats.Url, logger)
	err = natsClient.Connect()
	if err != nil {
		return err
	}

	// Subscribe to the provided subject
	err = natsClient.Subscribe(subject, func(msg *nats.Msg) {
		logger.Info("Received message", "subject", msg.Subject, "data", string(msg.Data))
	})
	if err != nil {
		return err
	}

	logger.Info("Subscription started", "subject", subject)

	// Wait for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Info("Context canceled, shutting down")
	case sig := <-signalChan:
		logger.Info("Received OS signal, shutting down", "signal", sig)
	}

	// Disconnect from NATS
	err = natsClient.Disconnect()
	if err != nil {
		logger.Error("Error during NATS disconnection", "error", err)
		return err
	}

	logger.Info("Graceful shutdown complete")
	return nil
}
