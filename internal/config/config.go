package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var cfg *Config

type Config struct {
	GRPC grpc `yaml:"grpc"`
}

type grpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
