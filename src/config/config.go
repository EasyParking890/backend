package config

import (
	"log"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	InternalPort string
	ExternalPort string
	RunMode      string
}

func GetConfig() *Config {
	//cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	//v, err := LoadConfig(cfgPath, "yml")

	//cfg, err := ParseConfig()
	var cfg Config
	envPort := "8001"
	if envPort != "" {
		cfg.Server.ExternalPort = envPort
		log.Printf("Set external port from environment -> %s", cfg.Server.ExternalPort)
	} else {
		cfg.Server.ExternalPort = cfg.Server.InternalPort
		log.Printf("Set external port from environment -> %s", cfg.Server.ExternalPort)
	}

	return &cfg
}

func ParseConfig() (*Config, error) {
	var cfg Config
	return &cfg, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/config/config-docker"
	} else if env == "production" {
		return "/config/config-production"
	} else {
		return "../config/config-development"
	}
}
