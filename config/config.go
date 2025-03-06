package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database `yaml:"database"`
}

type Database struct {
	DSN  string `yaml:"dsn"`
	Mock bool   `yaml:"mock"`
}

func ParseConfig() (*Config, error) {
	var cfg Config
	// TODO: check env to decide which file to parse
	err := cleanenv.ReadConfig("config/dev.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}
	return &cfg, nil
}
