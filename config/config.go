package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	GrpcPort string
	Smtp     Smtp
}

type Smtp struct {
	Sender   string
	Password string
}

func Load(path string) Config {
	err := godotenv.Load(path + "/.env") // load .env file if it exists
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		GrpcPort: conf.GetString("GRPC_PORT"),
		Smtp: Smtp{
			Sender:   conf.GetString("SMTP_SENDER"),
			Password: conf.GetString("SMTP_PASSWORD"),
		},
	}

	return cfg
}
