package config

import (
	"regexp"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Sites []Site
}

type Site struct {
	Name string
	Patterns []regexp.Regexp
	Url string
}

func GetConfig() (*Config, error) {
	var config Config
	configFile, err := GetConfigFile()
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
