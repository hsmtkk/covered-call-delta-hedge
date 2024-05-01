package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CallSymbol string
	PutSymbol  string
	Quantity   int
}

func (c Config) Save() error {
	f, err := os.Create("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewEncoder(f).Encode(c); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	return nil
}

func Load() (Config, error) {
	var c Config
	f, err := os.Open("config.yaml")
	if err != nil {
		return c, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		return c, fmt.Errorf("failed to decode YAML: %w", err)
	}
	return c, nil
}
