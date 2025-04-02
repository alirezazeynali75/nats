package config

type app struct {
	Env string `env:"ENV" envDefault:"development"`
	Name string `env:"NAME" envDefault:"nats-client"`
}