package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image       string            `yaml:"image"`
	Environment map[string]string `yaml:"environment"`

	Type    string `yaml:"type"`
	Command string `yaml:"command"`

	Ports   []string `yaml:"ports"`
	Volumes []string `yaml:"volumes"`
}

// findConfigFile walks up from the current directory looking for udock.yaml.
// Like git finding .git/. Works from any subdirectory of a project.
func findConfigFile() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := cwd
	for {
		candidate := filepath.Join(dir, "udock.yaml")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("no udock.yaml found in %s or any parent directory", cwd)
		}
		dir = parent
	}
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
