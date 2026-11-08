package config

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Sandbox SandboxConfig `yaml:"sandbox"`
	Mininet MininetConfig `yaml:"mininet"`
}
type SandboxConfig struct {
	Name string `yaml:"name"`
}
type MininetConfig struct {
	Topology   string `yaml:"topology"`
	Switch     string `yaml:"switch"`
	Controller string `yaml:"controller"`

	MAC     bool `yaml:"mac"`
	CLI     bool `yaml:"cli"`
	Cleanup bool `yaml:"cleanup"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config

	decoder := yaml.NewDecoder(
		bytes.NewReader(data),
	)

	decoder.KnownFields(true)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) Validate() error {
	if c.Sandbox.Name == "" {
		return fmt.Errorf("sandbox.name is required")
	}
	if c.Mininet.Topology == "" {
		return fmt.Errorf("sandbox.topology is required")
	}
	if c.Mininet.Switch == "" {
		return fmt.Errorf("sandbox.switch is required")
	}
	if c.Mininet.Controller == "" {
		return fmt.Errorf("sandbox.controller is required")
	}
	return nil
}
