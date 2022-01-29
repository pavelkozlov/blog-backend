package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Database Postgres `envPrefix:"POSTGRESQL_"`
}

type Postgres struct {
	Url      string `env:"URL,notEmpty"`
	User     string `env:"USER,notEmpty"`
	Password string `env:"PASSWORD,notEmpty"`
	Db       string `env:"DB,notEmpty"`
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
