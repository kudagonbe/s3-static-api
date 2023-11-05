package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Port      int    `env:"PORT" envDefault:"8080"`
	Endpoint  string `env:"STORAGE_ENDPOINT,required"`
	AccessKey string `env:"STORAGE_ACCESS_KEY,requird"`
	SecretKey string `env:"STORAGE_SECRET_KEY,required"`
	Bucket    string `env:"STORAGE_BUCKET,required"`
}

var c *Config

func Get() *Config {
	return c
}

func init() {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	c = &cfg
}
