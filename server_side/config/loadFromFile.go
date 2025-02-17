package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     int    `json:"port"`
	LogLevel string `json:"logLevel"`
}

func loadFromFile(fileName string) (Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
