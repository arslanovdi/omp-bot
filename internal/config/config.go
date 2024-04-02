package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

var cfg *Config

type Config struct {
	GRPC   grpc   `yaml:"grpc"`
	Jaeger Jaeger `yaml:"jaeger"`
}

type grpc struct {
	Host       string        `yaml:"host"`
	Port       string        `yaml:"port"`
	CtxTimeout time.Duration `yaml:"ctxTimeout"`
}

type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return fmt.Errorf("config.ReadConfigYML: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return fmt.Errorf("config.ReadConfigYML: %w", err)
	}

	return nil
}
