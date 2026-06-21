package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type Config struct {
	Env        string `yaml:"env" envDefault:"local"`
	DB         `yaml:"db"`
	HTTPServer `yaml:"http_server"`
	Auth       `yaml:"auth"`
}

type DB struct {
	Host     string `yaml:"host" envDefault:"localhost"`
	Port     int    `yaml:"port" envDefault:"5432"`
	Name     string `yaml:"name" envDefault:"notifier"`
	User     string `yaml:"user" envDefault:"root"`
	Password string `yaml:"password" envDefault:"root"`
	SSLMode  string `yaml:"ssl_mode" envDefault:"disable"`
}

type HTTPServer struct {
	Address      string `yaml:"address" envDefault:"0.0.0.0"`
	Timeout      string `yaml:"timeout" envDefault:"4s"`
	IddleTimeout string `yaml:"iddle_timeout" envDefault:"60s"`
}

type Auth struct {
	JWT `yaml:"jwt"`
}

type JWT struct {
	Secret string `yaml:"secret" envDefault:"secret"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
