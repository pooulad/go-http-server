package config

import (
	"github.com/caarlos0/env/v9"
)

type Server struct {
	Host string `env:"SERVERHOST"`
	Port int    `env:"SERVERPORT"`
}

type Config struct {
	Server
}

func LoadConfigOrPanic() Config {
	var config *Config = new(Config)
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}
	return *config
}
