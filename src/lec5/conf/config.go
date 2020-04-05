package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

type Config struct {
	Listen   string `env:"LISTEN" envDefault:"localhost:8080"`
	Greeting string `env:"GREETING" envDefault:"hello env"`
}

func Load() Config {
	return Config{
		Listen:   os.Getenv("LISTEN"),
		Greeting: os.Getenv("GREETING"),
	}
}

func Load2() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}
