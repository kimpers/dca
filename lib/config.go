package lib

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Coinbase struct {
		Secret       string `yaml:"secret"`
		Key          string `yaml:"key"`
		Passphrase   string `yaml:"passphrase"`
		BaseCurrency string `yaml:"baseCurrency"`
	}
	Pairs []struct {
		Name       string  `yaml:"name"`
		Percentage float64 `yaml:"percentage"`
	}
	Amount string `yaml:"amount"`
}

func (config *Config) ReadFile(path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Could not read config file. Error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("Failed to read yaml config file. Error: %v", err)
	}
}
