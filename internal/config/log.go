package config

type log struct {
	Level string `env:"LEVEL" envDefault:"info"`
}