package events

import (
	"log/slog"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	logger     *slog.Logger
	url        string
	connection *nats.Conn
}

func New(url string, logger *slog.Logger) *NatsClient {

	client := &NatsClient{
		logger:     logger,
		url:        url,
		connection: nil,
	}

	return client
}

func (c *NatsClient) Connect() error {
	c.logger.Info("Connecting to NATS server", "url", c.url)

	nc, err := nats.Connect(c.url)
	if err != nil {
		c.logger.Error("Failed to connect to NATS server", "error", err)
		return err
	}

	c.connection = nc
	c.logger.Info("Connected to NATS server")

	return nil
}

func (c *NatsClient) Disconnect() error {
	c.logger.Info("Disconnecting from NATS server")

	if c.connection != nil {
		err := c.connection.Drain()
		if err != nil {
			c.logger.Error("Failed to drain NATS connection", "error", err)
			return err
		}
		c.connection.Close()
		c.logger.Info("Disconnected from NATS server")
	} else {
		c.logger.Warn("No active connection to disconnect")
	}
	return nil
}

func (c *NatsClient) Publish(subject string, data []byte) error {
	if c.connection == nil {
		c.logger.Error("No active connection to publish message")
		return nats.ErrNoServers
	}

	c.logger.Info("Publishing message", "subject", subject)
	err := c.connection.Publish(subject, data)
	if err != nil {
		c.logger.Error("Failed to publish message", "error", err)
		return err
	}

	c.logger.Info("Message published successfully")
	return nil
}
func (c *NatsClient) Subscribe(subject string, callback nats.MsgHandler) error {
	if c.connection == nil {
		c.logger.Error("No active connection to subscribe to subject")
		return nats.ErrNoServers
	}

	c.logger.Info("Subscribing to subject", "subject", subject)
	subscription, err := c.connection.Subscribe(subject, callback)
	if err != nil {
		c.logger.Error("Failed to subscribe to subject", "error", err)
		return err
	}

	c.logger.Info("Subscribed to subject successfully", "subscription", subscription)
	return nil
}


