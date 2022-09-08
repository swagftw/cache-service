package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server *Server `yaml:"server"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Timeout  int    `yaml:"timeout"`
	Port     int    `yaml:"port"`
	Env      string `yaml:"env"`
	LogLevel string `yaml:"log_level"`
}

type Redis struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Password      string `yaml:"password"`
	DB            int    `yaml:"db"`
	TTL           int    `yaml:"ttl"`
	MaxMemorySize int    `yaml:"max_memory_size"`
	CacheStrategy string `yaml:"cache_strategy"`
}

// InitConfig initializes config
// this is temporary solution, in future use viper, and make getter setters over config struct
func InitConfig(configPath string) *Config {
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v\n", err)
	}

	cfg := new(Config)
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v\n", err)
	}

	return cfg
}
