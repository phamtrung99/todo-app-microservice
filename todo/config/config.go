package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var config *Config

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME"`
	Port        string `envconfig:"PORT"`
	Mysql       struct {
		Host     string `envconfig:"MYSQL_HOST"`
		Port     string `envconfig:"MYSQL_PORT"`
		User     string `envconfig:"MYSQL_ROOT_USER"`
		Password string `envconfig:"MYSQL_ROOT_PASSWORD"`
		Database string `envconfig:"MYSQL_DATABASE"`
	}
}

func LoadConfig() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return fmt.Errorf("failed to process config: %w", err)
	}
	config = &cfg
	return nil
}

func Get() *Config {
	if config == nil {
		panic("config not loaded: call LoadConfig() first")
	}
	return config
}
