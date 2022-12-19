package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	LogtailToken   LogtailToken 		`env:"LOGTAIL_TOKEN"`
	ProductionMode ProductionMode 	`env:"ENV"`

	ListenAddress	 string						`env:"LISTEN_ADDR" envDefault:"0.0.0.0"`
	ListenPort	 	 int							`env:"PORT" envDefault:"8080"`
}

func NewEnv() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewDotEnv() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return NewEnv()
}
