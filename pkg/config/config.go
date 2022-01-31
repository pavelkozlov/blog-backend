package config

import (
	"log"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Database Postgres `envPrefix:"POSTGRESQL_"`
	Server   Server   `envPrefix:"SERVER_"`
}

type Postgres struct {
	Url      string `env:"URL,notEmpty"`
	User     string `env:"USER,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
	Db       string `env:"DB,notEmpty"`
}

type Server struct {
	DebugMode bool `env:"DEBUG,notEmpty"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		log.Fatal("Error when config parse:", err.Error())
	}

	return cfg
}
