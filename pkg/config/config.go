package config

import "github.com/caarlos0/env/v9"

type Server struct {
	Host string `env:"SERVERHOST"`
	Port int    `env:"SERVERPORT"`
}

type Database struct {
	User string `env:"DBUSER"`
	Pass string `env:"DBPASS"`
	Host string `env:"DBHOST"`
	Port int    `env:"DBPORT"`
	Db   string `env:"DBDB"`
}

type Config struct {
	Server
	Database
}

func LoadConfigOrPanic() Config {

	var cfg *Config = new(Config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	return *cfg
}
