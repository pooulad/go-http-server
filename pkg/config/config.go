package config

import (
	"github.com/caarlos0/env/v9"
)

type Server struct {
	Host string `env:"SERVERHOST"`
	Port int    `env:"SERVERPORT"`
}

type Postgres struct {
	Username string `env:"PGUSER"`
	Password string    `env:"PGPASSWORD"`
	Port int    `env:"PGPORT"`
	Database string    `env:"PGDATABASE"`
	Host string    `env:"PGHOST"`
}

type Config struct {
	Server
	Postgres
}

func LoadConfigOrPanic() Config {
	var config *Config = new(Config)
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}
	return *config
}
