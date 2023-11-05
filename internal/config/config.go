package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port      int    `env:"PORT" envDefault:"8080"`
	Endpoint  string `env:"STORAGE_ENDPOINT,required"`
	AccessKey string `env:"STORAGE_ACCESS_KEY,required"`
	SecretKey string `env:"STORAGE_SECRET_KEY,required"`
	Bucket    string `env:"STORAGE_BUCKET,required"`
}

var c *Config

func Get() *Config {
	return c
}

func init() {
	if godotenv.Load() != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	c = &cfg
}
