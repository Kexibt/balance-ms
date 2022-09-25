package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const cfgFilename = "cfg.yml"

type Config struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	ExchangeRate    string        `yaml:"exchange_rate"`
	TimeoutConnPsql time.Duration `yaml:"timeout_conn_psql"`
	StringConnPsql  string        `yaml:"connection_psql"`
}

var cfg Config

func init() {
	update()
}

func update() {
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

func (c Config) GetConnectionTimeout() time.Duration {
	return c.TimeoutConnPsql
}

func (c Config) GetConnectionString() string {
	return c.StringConnPsql
}

func (c Config) GetExchangeRateLink() string {
	return c.ExchangeRate
}
