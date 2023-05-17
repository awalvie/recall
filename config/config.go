package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct holds the configuration for the application
type Config struct {
	Server struct {
		Name string `yaml:"name,omitempty"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

// Read reads a YAML config file and returns a Config struct
func Read(s string) (*Config, error) {
	// Read config file
	f, err := os.ReadFile(s)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	// Unmarshal config file
	var c Config
	if err := yaml.Unmarshal(f, &c); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return &c, nil
}
