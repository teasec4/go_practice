package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	// Provide a sensible default when CONFIG_PATH isn't set during local development.
	if configPath == "" {
		configPath = "config/local.yaml"
	}

	// check if file exists
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file does not exist or is not accessible (%s): %v", configPath, err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config from %s: %v", configPath, err)
	}

	return &cfg
}
