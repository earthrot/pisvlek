package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port     int `yaml:"port"`
	Database struct {
		Filename string `yaml:"filename"`
	} `yaml:"database"`
}

func New(filename string) (*Config, error) {
	var cfg Config

	f, err := os.Open(filename)
	if err != nil {
		return &cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil

}
