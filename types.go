package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type Program struct {
	id         int
	name       string
	platform   string
	submission string
	bounty     bool
}

type Target struct {
	id         int
	name       string
	category   string
	scope      bool
	program_id int
}

type Bugcrowd struct {
}

func loadConfig() Config {
	var config Config

	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return config
	}

	// Unmarshal YAML into the Config struct
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return config
	}

	return config
}
