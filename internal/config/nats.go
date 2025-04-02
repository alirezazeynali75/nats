package config

type nats struct {
	Url string `env:"URL" envDefault:"nats://localhost:4222"`
}