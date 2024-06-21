package config

import (
	"os"

	"github.com/PabloBagliere/B-21/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	NameSystem string                      `yaml:"nameSystem"`
	Version    string                      `yaml:"version"`
	Services   map[interface{}]interface{} `yaml:"Services"`
}

// ReadFile reads the configuration file
func ReadFile() (*Config, error) {
	var cfg Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, errors.NewConfigError("Error reading the configuration file", err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, errors.NewConfigError("Error unmarshalling the configuration file", err)
	}

	return &cfg, nil
}
