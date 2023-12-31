package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	Endpoint     string `env:"STORAGE_ENDPOINT,required"`
	AccessKey    string `env:"STORAGE_ACCESS_KEY,required"`
	SecretKey    string `env:"STORAGE_SECRET_KEY,required"`
	Bucket       string `env:"STORAGE_BUCKET,required"`
	UsePathStyle bool   `env:"STORAGE_USE_PATH_STYLE" envDefault:"false"`
	AwsConfig    aws.Config
}

var c *Config

func Get() *Config {
	return c
}

func init() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	if godotenv.Load(envFile) != nil {
		fmt.Printf("No %s file found\n", envFile)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	awsConfig := *aws.NewConfig()
	awsConfig.Region = "ap-northeast-1"
	awsConfig.BaseEndpoint = aws.String(cfg.Endpoint)
	awsConfig.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     cfg.AccessKey,
			SecretAccessKey: cfg.SecretKey,
		}, nil
	})

	cfg.AwsConfig = awsConfig

	c = &cfg
}
