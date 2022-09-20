package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const cfgFilename = "cfg.yml"

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var cfg Config

func init() {
	b, err := os.ReadFile(cfgFilename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() Config {
	return cfg
}

func (c Config) GetHostPort() string {
	return c.Host + ":" + c.Port
}
