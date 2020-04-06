package main

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Listen   string `env:"LISTEN" envDefault:"localhost:8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

func loadConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, env.Parse(cfg)
}
