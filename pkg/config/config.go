package config

import "github.com/caarlos0/env/v9"

type Server struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}
type Config struct {
	Server
}

func LoadConfigOrPanic() Config {
	var config *Config = new(Config)

	if err := env.Parse(config); err != nil {
		panic(err)
	}

	return *config
}
